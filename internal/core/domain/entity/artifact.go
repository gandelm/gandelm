package entity

type Artifact struct {
	Name    string
	Type    ArtifactType
	Version string
	URL     string
}

type Artifacts []*Artifact

func (a Artifacts) ToTags() ArtifactTag {
	tags := make(ArtifactTag)
	for _, artifact := range a {
		tags[artifact.Name] = artifact.Name
	}

	return tags
}

type ArtifactTag map[string]string
