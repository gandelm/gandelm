package provider

import (
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/gandelm/gandelm/internal/core/infrastructure/etcd"
	"github.com/gandelm/gandelm/internal/core/usecase"
)

func NewDeploymentRORepository(container container.Containerer) repository.DeploymentRORepository {
	return etcd.NewDeploymentRORepository(container)
}

func NewCatalogCreator(container container.Containerer) usecase.CatalogCreator {
	return usecase.NewCatalogCreate(
		NewCatalogRWRepository(container),
	)
}

func NewCatalogUpdator(container container.Containerer) usecase.CatalogUpdator {
	return usecase.NewCatalogUpdate(
		NewCatalogRWRepository(container),
	)
}

func NewCatalogRORepository(container container.Containerer) repository.CatalogRORepository {
	return etcd.NewCatalogRORepository(container)
}

func NewCatalogRWRepository(container container.Containerer) repository.CatalogRWRepository {
	return etcd.NewCatalogRepository(container)
}

func NewWorkloadRORepository(container container.Containerer) repository.WorkloadRORepository {
	return etcd.NewWorkloadRORepository(container)
}

func NewLabelRORepository(container container.Containerer) repository.LabelRORepository {
	return etcd.NewLabelRORepository(container)
}

func NewLabelRWRepository(container container.Containerer) repository.LabelRWRepository {
	return etcd.NewLabelRepository(container)
}

func NewVersionRORepository(container container.Containerer) repository.VersionRORepository {
	return etcd.NewVersionRORepository(container)
}

func NewVersionRWRepository(container container.Containerer) repository.VersionRWRepository {
	return etcd.NewVersionRepository(container)
}
