package src

type Permission struct{
	name string
}

func (r *Rbac)NewPermission(name string) *Permission {
	new_permission := &Permission{name: name}
	r.permissions = append(r.permissions, new_permission)
	return new_permission
}