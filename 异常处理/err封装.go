package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// RequeueError represents an requeueError resource error.
type RequeueError struct {
	error
}

// NewRequeueErr creates an requeueError resource error.
func NewRequeueErr(err error) error {
	return RequeueError{err}
}

// IsRequeueErr returns true if the err is RequeueError error type.
func IsRequeueErr(err error) bool {
	_, ok := err.(RequeueError)
	return ok
}

func main() {
	err := NewRequeueErr(errors.New("hyqtest"))
	if IsRequeueErr(err) {
		fmt.Println("IsRequeueErr")
	}

	errNew := errors.Wrap(err, "new err")
	fmt.Println(errNew.Error())
	if IsRequeueErr(errNew) {
		fmt.Println("IsRequeueErr new")
	}
}
