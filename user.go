package src

type User struct {
	id    int64
	name  string
	//roles []*Role
	roleMap map[string]*Role
}

func (r *Rbac)NewUser(id int64, name string, roles []*Role) *User {
	new_user := &User{id: id, name: name}
	new_user.roleMap = map[string]*Role{}
	for _, role := range roles {
		new_user.roleMap[role.name] = role
	}
	r.users = append(r.users, new_user)
	return new_user
}

/*func (user *User) loads () error {
}*/

func (user *User) grant(roles []*Role) {
	//user.roles = roles
	for _, role := range roles {
		if _, ok := user.roleMap[role.name]; ok {
			continue
		}
		user.roleMap[role.name] = role
	}
	return
}

func (user *User) permit(permission *Permission) bool {
	for _, role := range user.roleMap {
		if ok := role.permit(permission); ok {
			return true
		}
	}
	return false
}

func (user *User) revoke(revoke_role *Role) {
	if _, ok := user.roleMap[revoke_role.name]; !ok {
		return
	}
	delete(user.roleMap, revoke_role.name)
	return
}
