package gandelmcatalog

type Phase = string

const (
	NONE        Phase = ""
	INITIALIZED Phase = "Initialized"
	UPDATED     Phase = "Updated"
	TERMINATED  Phase = "Terminated"
)
