package src

import (
	"fmt"
	"testing"
)

func TestNewRole(t *testing.T) {
	p0 := NewPermission("p0")
	p1 := NewPermission("p1")
	ps := []*Permission{p0, p1}
	r0 := NewRole(ps)
	fmt.Println("r0: ", r0)
}