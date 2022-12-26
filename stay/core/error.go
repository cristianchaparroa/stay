package core

type ParamError struct {
	// Name of the field with issues
	Name string `json:"name"`

	// Reason why the field has an issue
	Reason string `json:"reason"`
}

type Error struct {
	// Message error
	Message string `json:"message"`

	// Status code
	Status int `json:"status"`

	// InvalidParams is used when a request needs a
	// params and those are missing
	InvalidParams []*ParamError `json:"invalid-params,omitempty"`
}

func NewError(message string, status int) *Error {
	return &Error{
		Message: message,
		Status:  status,
	}
}
