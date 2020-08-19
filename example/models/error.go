//define the sington error
package models

import (
	"strconv"
)

type singtonError struct {
	msg  string
	code uint32
}

func NewSingtonError(code uint32, text string) error {
	return &singtonError{text, code}
}

func (e *singtonError) Error() string {
	return "Error Code: " + strconv.FormatUint(uint64(e.code), 10) + "; Message: " + e.msg
}
