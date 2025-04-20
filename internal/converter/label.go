package converter

import (
	v1 "github.com/gandelm/gandelm/api/v1"
	labelv1 "github.com/gandelm/gandelm/generated/protocol/label/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

func MakeLabelPb(label *entity.Label) *labelv1.Label {
	return &labelv1.Label{
		Id:          label.ID,
		Title:       label.Title,
		Description: label.Description,
		Color:       label.Color,
	}
}

func MakeLabelsPb(labels entity.Labels) []*labelv1.Label {
	results := make([]*labelv1.Label, 0, len(labels))
	for _, label := range labels {
		results = append(results, MakeLabelPb(label))
	}

	return results
}

func MakeLabel(label *v1.GandelmLabel) *entity.Label {
	return &entity.Label{
		ID:          label.Name,
		Title:       label.Spec.Title,
		Description: label.Spec.Description,
		Color:       label.Spec.Color,
	}
}

func MakeLabels(labelList *v1.GandelmLabelList) entity.Labels {
	results := make(entity.Labels, 0, len(labelList.Items))
	for _, label := range labelList.Items {
		results = append(results, MakeLabel(&label))
	}

	return results
}
