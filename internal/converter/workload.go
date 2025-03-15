package converter

import (
	v1 "github.com/gandelm/gandelm/api/v1"
	workloadv1 "github.com/gandelm/gandelm/generated/protocol/workload/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

func MakeExternalLinksPb(links entity.ExternalLinks) []*workloadv1.ExternalLink {
	results := make([]*workloadv1.ExternalLink, 0, len(links))
	for _, link := range links {
		results = append(results, &workloadv1.ExternalLink{
			Title: link.Title,
			Url:   link.URL,
		})
	}

	return results
}

func MakeWorkloadPb(workload *entity.Workload) *workloadv1.Workload {
	return &workloadv1.Workload{
		WorkloadId:    workload.ID,
		Endpoint:      workload.Endpoint,
		Entrypoint:    workload.Entrypoint,
		ExternalLinks: MakeExternalLinksPb(workload.ExternalLinks),
	}
}

func MakeWorkloadsPb(workloads entity.Workloads) []*workloadv1.Workload {
	results := make([]*workloadv1.Workload, 0, len(workloads))
	for _, workload := range workloads {
		results = append(results, MakeWorkloadPb(workload))
	}

	return results
}

func MakeWorkload(workload *v1.Gandelm) *entity.Workload {
	externalLinks := make([]*entity.ExternalLink, 0, len(workload.Spec.ExternalLinks))
	for _, link := range workload.Spec.ExternalLinks {
		externalLinks = append(externalLinks, &entity.ExternalLink{
			Title: link.Title,
			URL:   link.URL,
		})
	}

	return &entity.Workload{
		ID:            workload.Name,
		Endpoint:      workload.Spec.Endpoint,
		Entrypoint:    workload.Spec.Entrypoint,
		ExternalLinks: externalLinks,
	}
}

func MakeWorkloads(workload []v1.Gandelm) entity.Workloads {
	results := make(entity.Workloads, 0, len(workload))
	for _, w := range workload {
		results = append(results, MakeWorkload(&w))
	}

	return results
}
