package src

type Permission struct{
	name string
}

// create a new permission
func (r *Rbac)NewPermission(name string) *Permission {
	new_permission := &Permission{name: name}
	r.permissions = append(r.permissions, new_permission)
	return new_permission
}