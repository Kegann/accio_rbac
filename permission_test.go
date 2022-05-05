package src

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermission_New(t *testing.T) {
	rr0 := Rbac{}
	p0 := rr0.NewPermission("t0")
	p1 := rr0.NewPermission("t0")
	fmt.Println("new permission: ", p0, p1, p0 == p1, reflect.DeepEqual(p0, p1))
}