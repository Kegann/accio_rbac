package src

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	rr0 := Rbac{users: map[int64]*User{}, roles: map[string]*Role{}, permissions: map[string]*Permission{}}
	p0 := rr0.NewPermission("p0")
	p1 := rr0.NewPermission("p1")
	ps0 := []*Permission{p0}
	ps1 := []*Permission{p1}
	r0 := rr0.NewRole("r0", ps0)
	r1 := rr0.NewRole("r1", ps1)
	rs := []*Role{r0, r1}
	// c
	u1 := rr0.NewUser(1, "u1", nil)
	u2str := `{
		"userid": 2,
		"username": "u2"
	}`
	u2err, u2 := rr0.LoadUser(u2str)
	// r
	u1.grant(rs)
	fmt.Println("c u1: ", u1, u1.permit(p1))
	if (u2err == nil) {
		fmt.Println("c u2: ", u2, u2.permit(p1))
	} else {
		fmt.Println("c u2 err: ", u2err)
	}
	for _, role := range u1.RoleMap {
		fmt.Println("r u1 role ", role.Name)
	}
	// u d
	// u2.grant(rs)
	for _, role := range u1.RoleMap {
		fmt.Println("r u u1 role ", role.Name)
	}
	r2 := rr0.NewRole("r0", ps0)
	u2.revoke(r2)
	for _, role := range u1.RoleMap {
		fmt.Println("r d u1 role ", role.Name)
	}
	fmt.Println("last,", u1.RoleMap)
	fmt.Println("u0: ", u1, u1.permit(p1))
}
