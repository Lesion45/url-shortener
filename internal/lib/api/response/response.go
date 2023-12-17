package response

import (
	"fmt"
	"github.com/go-playground/validator"
	"strings"
)

type Response struct {
	Status string `json: "status"`
	Error  string `json: "error, omitempty"`
}

const (
	StatusOk    = "Ok"
	StatusError = "Error"
)

func OK() Response {
	return Response{Status: StatusOk}
}

func Error(msg string) Response {
	return Response{Status: StatusError,
		Error: msg}
}

func ValidataionError(errs validator.ValidationErrors) Response {
	var errsMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is a required field", err.Field()))

		case "url":
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is not a valid url", err.Field()))

		default:
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is not valid"))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errsMsgs, ", "),
	}
}
