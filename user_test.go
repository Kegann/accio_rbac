package src

import "testing"

func TestNewUser(t *testing.T) {
	p0 := NewPermission("p0")
	p1 := NewPermission("p1")
	ps := []*Permission{p0, p1}
	r0 := NewRole(ps)
	u0 := NewUser(1)
}
