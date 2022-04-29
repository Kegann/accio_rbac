package src

type Role struct {
	name        string
	permissions []*Permission
}

func NewRole(name string, permissions []*Permission) *Role {
	return &Role{name: name, permissions: permissions}
}

func (role *Role) permit(comp_permission *Permission) bool {
	for _, permission := range role.permissions {
		if comp_permission == permission {
			return true
		}
	}
	return false
}

func (role *Role) assign(new_permission *Permission) {
	// check if exist
	// todo optimize
	for _, permission := range role.permissions {
		if new_permission.name == permission.name {
			return
		}
	}
	role.permissions = append(role.permissions, new_permission)
	return
}

func (role *Role) revoke(revoke_permission *Permission) {
	// check if exist before delete
	// todo optimize
	for i := 0; i < len(role.permissions); i++ {
		if revoke_permission.name == role.permissions[i].name {
			role.permissions = append(role.permissions[:i], role.permissions[i+1:]...)
			return
		}
	}
	return
}
