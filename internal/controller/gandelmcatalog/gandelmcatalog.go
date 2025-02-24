package gandelmcatalog

import (
	"context"
	"fmt"
	"time"

	gandelmcomv1 "github.com/gandelm/gandelm/api/v1"
	v1 "github.com/gandelm/gandelm/api/v1"
	"github.com/gandelm/gandelm/internal/container"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

var finalizer = "gandelm.com/finalizer"

// GandelmCatalogReconciler reconciles a GandelmCatalog object
type GandelmCatalogReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	Container container.Containerer
}

// +kubebuilder:rbac:groups=gandelm.com,resources=gandelmcatalogs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gandelm.com,resources=gandelmcatalogs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=gandelm.com,resources=gandelmcatalogs/finalizers,verbs=update
func (r *GandelmCatalogReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var catalog v1.GandelmCatalog
	if err := r.Get(ctx, req.NamespacedName, &catalog); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		logger.Error(err, "unable to fetch GandelmCatalog")
		return ctrl.Result{RequeueAfter: time.Duration(30 * time.Second)}, err
	}

	if !catalog.DeletionTimestamp.IsZero() {
		return r.Delete(ctx, catalog)
	}

	if catalog.Status.Phase == NONE {
		return r.Create(ctx, catalog)
	}

	return r.Update(ctx, catalog)
}

// SetupWithManager sets up the controller with the Manager.
func (r *GandelmCatalogReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gandelmcomv1.GandelmCatalog{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Named("gandelmcatalog").
		Complete(r)
}

func (r *GandelmCatalogReconciler) Create(ctx context.Context, catalog v1.GandelmCatalog) (ctrl.Result, error) {
	if err := r.Container.Command().HelmInstall(catalog.Name, "./repositories/gandelm/manifests/nginx"); err != nil {
		return ctrl.Result{RequeueAfter: time.Duration(time.Minute)}, nil
	}

	// フック結果でステータスを更新
	catalog.Status.Phase = INITIALIZED
	catalog.Status.Message = "初期化スクリプトをフックしました"
	catalog.Status.Timestamp = metav1.Time{Time: time.Now()}

	if err := r.Status().Update(ctx, &catalog); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *GandelmCatalogReconciler) Update(ctx context.Context, catalog v1.GandelmCatalog) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	if catalog.ObjectMeta.Generation == catalog.Status.ObservedGeneration {
		logger.Info(fmt.Sprintf("Synced: %s", catalog.Name))
		return ctrl.Result{}, nil
	}

	// if _, err := r.Container.Github().HookAction(ctx, "test", map[string]string{
	// 	"name": catalog.Name,
	// 	"type": "update",
	// }); err != nil {
	// 	return ctrl.Result{
	// 		RequeueAfter: time.Duration(time.Minute),
	// 	}, err
	// }

	catalog.Status.Phase = UPDATED
	catalog.Status.Message = "更新スクリプトをフックしました"
	catalog.Status.ObservedGeneration = catalog.ObjectMeta.Generation
	catalog.Status.Timestamp = metav1.Time{Time: time.Now()}
	if err := r.Status().Update(ctx, &catalog); err != nil {
		logger.Error(err, "unable to update GandelmCatalog status")
		return ctrl.Result{
			RequeueAfter: time.Duration(time.Minute),
		}, err
	}

	return ctrl.Result{}, nil
}

func (r *GandelmCatalogReconciler) Delete(ctx context.Context, catalog v1.GandelmCatalog) (ctrl.Result, error) {
	if len(catalog.Finalizers) == 0 {
		return ctrl.Result{}, nil
	}

	for _, f := range catalog.Finalizers {
		if f != finalizer {
			continue
		}

		if err := r.Container.Command().HelmUnInstall(catalog.Name); err != nil {
			return ctrl.Result{RequeueAfter: time.Duration(time.Minute)}, err
		}
	}

	catalog.Finalizers = []string{}
	if err := r.Client.Update(ctx, &catalog); err != nil {
		return ctrl.Result{RequeueAfter: time.Duration(time.Minute)}, err
	}

	return ctrl.Result{}, nil
}
