package errors

type ProblemDetails struct {
	Type          string         `json:"type"`
	Title         string         `json:"title"`
	Detail        string         `json:"detail,omitempty"`
	Details       string         `json:"details,omitempty"`
	InvalidParams []InvalidParam `json:"invalid_params,omitempty"`
	Status        int            `json:"status_code,omitempty"`
}

type InvalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func FromErrorToProblemDetails(err error) *ProblemDetails {
	return &ProblemDetails{}
}
