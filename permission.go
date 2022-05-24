package src

type Permission struct{
	// todo autoincr id?
	Name string
}

// create a new permission
func (r *Rbac)NewPermission(name string) *Permission {
	if _, ok := r.permissions[name]; ok {
		return nil
	}
	new_permission := &Permission{Name: name}
	r.permissions[name] = new_permission
	return new_permission
}