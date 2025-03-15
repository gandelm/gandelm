package entity

import (
	"strings"
)

func NewWorkloadID(version, name string) WorkloadID {
	id := strings.ReplaceAll(version, ".", "-") + "-" + name
	return strings.ToLower(id)
}

type WorkloadID = string
type Workloads []*Workload
type Workload struct {
	ID            WorkloadID
	Endpoint      string
	Entrypoint    string
	ExternalLinks []*ExternalLink

	Artifacts Artifacts
}

func (w *Workload) Namespace() string {
	return w.ID
}

type ExternalLinks []*ExternalLink
type ExternalLink struct {
	Title string
	URL   string
}
