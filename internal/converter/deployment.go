package converter

import (
	workloadv1 "github.com/gandelm/gandelm/generated/protocol/workload/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func MakeDeploymentsPb(deployments []appsv1.Deployment) []*workloadv1.Deployment {
	results := make([]*workloadv1.Deployment, 0, len(deployments))
	for _, deployment := range deployments {
		results = append(results, MakeDeploymentPb(deployment))
	}

	return results
}

func MakeDeploymentPb(deployment appsv1.Deployment) *workloadv1.Deployment {
	message := ""
	if len(deployment.Status.Conditions) > 0 {
		message = deployment.Status.Conditions[0].Message
	}

	return &workloadv1.Deployment{
		Containers:    MakeContainers(deployment.Spec.Template.Spec.Containers),
		ReplicaStatus: MakeReplicaStatus(deployment.Status),
		Message:       message,
	}
}

func MakeContainers(containers []corev1.Container) []*workloadv1.Container {
	results := make([]*workloadv1.Container, 0, len(containers))
	for _, container := range containers {
		results = append(results, MakeContainer(container))
	}

	return results
}

func MakeContainer(container corev1.Container) *workloadv1.Container {
	return &workloadv1.Container{
		Name:   container.Name,
		Image:  container.Image,
		IsInit: false,
	}
}

func MakeReplicaStatus(status appsv1.DeploymentStatus) *workloadv1.ReplicaStatus {
	return &workloadv1.ReplicaStatus{
		Desired:   uint32(status.Replicas),
		Ready:     uint32(status.ReadyReplicas),
		Available: uint32(status.AvailableReplicas),
		Updated:   uint32(status.UpdatedReplicas),
	}
}
