package src

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	rr0 := Rbac{}
	p0 := rr0.NewPermission("p0")
	p1 := rr0.NewPermission("p1")
	ps0 := []*Permission{p0}
	ps1 := []*Permission{p1}
	r0 := rr0.NewRole("r0", ps0)
	r1 := rr0.NewRole("r1", ps1)
	rs := []*Role{r0, r1}
	// c
	u1 := rr0.NewUser(1, "u1", nil)
	// r
	fmt.Println("c u1: ", u1, u1.permit(p1))
	for _, role := range u1.roles {
		fmt.Println("r u1 role ", role.name)
	}
	// u d
	u1.grant(rs)
	for _, role := range u1.roles {
		fmt.Println("r u u1 role ", role.name)
	}
	r2 := rr0.NewRole("r0", ps0)
	u1.revoke(r2)
	for _, role := range u1.roles {
		fmt.Println("r d u1 role ", role.name)
	}
	fmt.Println("last,", u1.roles)
	fmt.Println("u0: ", u1, u1.permit(p1))
}
