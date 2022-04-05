package crosscutting

import (
	"fmt"
)

type StandardError struct {
	code    string
	title   string
	message string
	details string
}

func NewStdError(code string, title string, msg string, details string) *StandardError {
	return &StandardError{
		code:    code,
		title:   title,
		message: msg,
		details: details,
	}
}

func (stdErr *StandardError) Error() string {
	return fmt.Sprintf("code: %v, title: %v, message: %v, details: %v", stdErr.code, stdErr.title, stdErr.message, stdErr.details)
}

type ValidationErrors struct {
	Errors []*ValidationError
}

type ValidationError struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Errors: make([]*ValidationError, 0),
	}
}

func (verrs *ValidationErrors) Add(name string, reason string) {
	verrs.Errors = append(verrs.Errors, &ValidationError{
		Name:   name,
		Reason: reason,
	})
}
