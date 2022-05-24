package src

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermission_New(t *testing.T) {
	rr0 := Rbac{users: map[int64]*User{}, roles: map[string]*Role{}, permissions: map[string]*Permission{}}
	p0 := rr0.NewPermission("t0")
	p1 := rr0.NewPermission("t0")
	p2str := `{
 		"name": "p2"
	}`
	err, p2 := rr0.LoadPermission(p2str)
	if err == nil {
		fmt.Println("p2: ", p2)
	} else {
		fmt.Println(err)
	}
	fmt.Println("new permission: ", p0, p1, p0 == p1, reflect.DeepEqual(p0, p1))
}