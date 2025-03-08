package entity

import (
	"strings"
	"time"
)

func NewWorkloadID(version, name string) WorkloadID {
	id := strings.ReplaceAll(version, ".", "-") + "-" + name
	return strings.ToLower(id)
}

type WorkloadID = string
type Workloads []*Workload
type Workload struct {
	ID        WorkloadID
	CreatedBy User
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w *Workload) Namespace() string {
	return w.ID
}
