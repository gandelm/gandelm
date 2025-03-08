package entity

import (
	"strings"
	"time"
)

type Workloads []*Workload
type Workload struct {
	ID        WorkloadID
	CreatedBy User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WorkloadID = string

func NewWorkloadID(version, name string) WorkloadID {
	id := strings.ReplaceAll(version, ".", "-") + "-" + name
	return strings.ToLower(id)
}
