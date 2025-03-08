package entity

type Artifact struct {
	Name string
	Type ArtifactType
}

type Artifacts []*Artifact

func (a Artifacts) ToTags() ArtifactTag {
	tags := make(ArtifactTag)
	for _, artifact := range a {
		tags[artifact.Name] = artifact.Type
	}

	return tags
}

type ArtifactTag map[string]string
