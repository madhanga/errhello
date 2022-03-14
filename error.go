package errhello

import (
	"bytes"
	"fmt"
)

type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s:", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code and message
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, " <%s> ", e.Code)
		}
		if e.Message != "" {
			buf.WriteString("Message: " + e.Message)
		}
	}

	return buf.String()
}

func ErrorCode(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}

	return EINTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}

	return "An internal error has occurred. Please contact support"
}
