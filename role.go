package src

type Role struct{
	permissions []*Permission
}

func NewRole(permissions []*Permission) *Role {
	return &Role{permissions:permissions}
}

func (role *Role) permit(comp_permission *Permission) bool {
	for _, permission := range role.permissions {
		if comp_permission == permission {
			return true
		}
	}
	return false
}