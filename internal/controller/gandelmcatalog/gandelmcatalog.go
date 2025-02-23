package gandelmcatalog

import (
	"context"
	"fmt"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	gandelmcomv1 "github.com/gandelm/gandelm/api/v1"
	v1 "github.com/gandelm/gandelm/api/v1"
)

// GandelmCatalogReconciler reconciles a GandelmCatalog object
type GandelmCatalogReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=gandelm.com,resources=gandelmcatalogs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gandelm.com,resources=gandelmcatalogs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=gandelm.com,resources=gandelmcatalogs/finalizers,verbs=update
func (r *GandelmCatalogReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var catalog v1.GandelmCatalog
	err := r.Get(ctx, req.NamespacedName, &catalog)
	if apierrors.IsNotFound(err) {
		return r.Delete(ctx, req.NamespacedName)
	} else if err != nil {
		logger.Error(err, "unable to fetch GandelmCatalog")
		return ctrl.Result{
			RequeueAfter: time.Duration(time.Minute),
		}, err
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
	logger := log.FromContext(ctx)
	logger.Info(fmt.Sprintf("新しい GandelmCatalog を検知しました(%s)", catalog.Name))
	logger.Info("Github Actions をフックします")

	// ここでフックする

	// フック結果でステータスを更新
	catalog.Status.Phase = INITIALIZED
	catalog.Status.Message = "初期化スクリプトをフックしました"
	catalog.Status.Timestamp = metav1.Time{Time: time.Now()}
	if err := r.Status().Update(ctx, &catalog); err != nil {
		logger.Error(err, "unable to update GandelmCatalog status")
		return ctrl.Result{
			RequeueAfter: time.Duration(time.Minute),
		}, err
	}

	return ctrl.Result{}, nil
}

func (r *GandelmCatalogReconciler) Update(ctx context.Context, catalog v1.GandelmCatalog) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	catalog.Status.Phase = UPDATED
	catalog.Status.Message = "更新スクリプトをフックしました"
	catalog.Status.Timestamp = metav1.Time{Time: time.Now()}

	logger.Info(fmt.Sprintf("GandelmCatalog を更新します(%s)", catalog.Name))
	if err := r.Status().Update(ctx, &catalog); err != nil {
		logger.Error(err, "unable to update GandelmCatalog status")
		return ctrl.Result{
			RequeueAfter: time.Duration(time.Minute),
		}, err
	}

	return ctrl.Result{}, nil
}

func (r *GandelmCatalogReconciler) Delete(ctx context.Context, req types.NamespacedName) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info(fmt.Sprintf("GandelmCatalog が削除されました(%s)", req.Name))
	logger.Info("Github Actins をフックします")
	return ctrl.Result{}, nil
}
