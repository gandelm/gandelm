package gandelmcatalog

type Phase = string

const (
	NONE           Phase = ""
	ADDEDFINALIZER Phase = "AddedFinalizer"
	INITIALIZED    Phase = "Initialized"
	UPDATED        Phase = "Updated"
	TERMINATED     Phase = "Terminated"
)
