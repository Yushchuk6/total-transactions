package entity

import (
	"errors"
)

var (
	//ErrVariableOutOfRange Error: variable out of allowed range
	ErrVariableOutOfRange = errors.New("Error: variable out of allowed range")
	//ErrIncorrectValue Error: passed incorrect value
	ErrIncorrectValue = errors.New("Error: passed incorrect value")
	//ErrResponveServerNOTOK Error: got NOTOK for API server
	ErrResponveServerNOTOK = errors.New("Error: got NOTOK for API server")
)