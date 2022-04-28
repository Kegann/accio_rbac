package src

type User struct {
	id int64
	roles []*Role
}

func NewUser (id int64) *User {
	return &User{id:id}
}

/*func (user *User) loads () error {
}*/

func (user *User) grant (roles []*Role) {
	user.roles = roles
	return
}

func (user *User) permit (permission *Permission) bool {
	for _, role := range user.roles {
		if ok := role.permit(permission); ok {
			return true
		}
	}
	return false
}
