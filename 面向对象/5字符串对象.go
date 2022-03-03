package main

import (
	"fmt"
	"regexp"
)

// ProcedureType represent type of procedure.
type ProcedureType string

const (
	// RoleCheckType represents Role check procedure
	RoleCheckType ProcedureType = "RoleCheck"
)

func main() {

	if matches := regexp.MustCompile(`release ([^ ]*)`).FindStringSubmatch("Kylin Linux Advanced Server release V10 ( Sword)"); matches != nil {
		fmt.Println(matches[1])
	}

	fmt.Println(testType(RoleCheckType))
	fmt.Println(testType("hyq"))
}

func testType(t ProcedureType) string {
	fmt.Println(t)
	return string(t)
}
