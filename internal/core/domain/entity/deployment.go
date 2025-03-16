package entity

// Deployment は Kubernetes 上のデプロイメントリソースを表すエンティティです。
type Deployments []*Catalog
type Deployment struct {
	Name          string
	Containers    []*Container
	ReplicaStatus *ReplicaStatus
	Message       string
}

type Container struct {
	Name            string
	Image           string
	IsInitContainer bool
}

type ReplicaStatus struct {
	Replicas          int32
	ReadyReplicas     int32
	AvailableReplicas int32
	UpdatedReplicas   int32
}
