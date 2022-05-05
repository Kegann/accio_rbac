package src

type User struct {
	id    int64
	name  string
	roles []*Role
}

func (r *Rbac)NewUser(id int64, name string, roles []*Role) *User {
	new_user := &User{id: id, name: name, roles: roles}
	r.users = append(r.users, new_user)
	return new_user
}

/*func (user *User) loads () error {
}*/

func (user *User) grant(roles []*Role) {
	user.roles = roles
	return
}

func (user *User) permit(permission *Permission) bool {
	for _, role := range user.roles {
		if ok := role.permit(permission); ok {
			return true
		}
	}
	return false
}

func (user *User) revoke(revoke_role *Role) {
	// todo optimize hash+link
	for i := 0; i < len(user.roles); i++ {
		if revoke_role.name == user.roles[i].name {
			user.roles = append(user.roles[:i], user.roles[i+1:]...)
			return
		}
	}
	return
}
