package src

import (
	"encoding/json"
)

type User struct {
	Id    int64 `json:"userid"`
	Name  string `json:"username"`
	//roles []*Role
	RoleMap map[string]*Role
}

// create a new user
func (r *Rbac)NewUser(id int64, name string, roles []*Role) *User {
	if _, ok := r.users[id]; ok {
		return nil
	}
	new_user := &User{Id: id, Name: name}
	new_user.RoleMap = map[string]*Role{}
	for _, role := range roles {
		new_user.RoleMap[role.Name] = role
	}
	r.users[id] = new_user
	return new_user
}

func (r *Rbac) loads (src string) error {
	u := new(User)
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		return err
	}
	r.users[u.Id] = u
	return nil
}


func (user *User) grant(roles []*Role) {
	//user.roles = roles
	for _, role := range roles {
		if _, ok := user.RoleMap[role.Name]; ok {
			continue
		}
		user.RoleMap[role.Name] = role
	}
	return
}

// check if user has the permission
func (user *User) permit(permission *Permission) bool {
	for _, role := range user.RoleMap {
		if ok := role.permit(permission); ok {
			return true
		}
	}
	return false
}

// delete user's role
func (user *User) revoke(revoke_role *Role) {
	if _, ok := user.RoleMap[revoke_role.Name]; !ok {
		return
	}
	delete(user.RoleMap, revoke_role.Name)
	return
}
