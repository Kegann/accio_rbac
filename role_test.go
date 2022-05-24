package src

import (
	"fmt"
	"testing"
)

func TestNewRole(t *testing.T) {
	rr0 := Rbac{users: map[int64]*User{}, roles: map[string]*Role{}, permissions: map[string]*Permission{}}
	p0 := rr0.NewPermission("p0")
	p1 := rr0.NewPermission("p1")
	ps := []*Permission{p0, p1}
	// c
	r0 := rr0.NewRole("r0", ps)
	fmt.Println("r0: ", r0)
	// r
	//permissions := r0.permissions
	//for idx, permission := range permissions {
	//	fmt.Println("permission: ", idx, permission.name)
	//}
	//fmt.Println("permissions: ", permissions)
	// u d
	p2 := rr0.NewPermission("p2")
	p3 := rr0.NewPermission("p3")
	r0.assign(p2)
	fmt.Println("r0: ", r0.PermissionMap, p2, p3)
	r0.revoke(p3)
	fmt.Println("r0: ", r0.PermissionMap)
	//for idx, permission := range r0.permissions {
	//	fmt.Println("u permission: ", idx, permission.name)
	//}
}
