package src

import "encoding/json"

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

// load permission from json str
func (r *Rbac) LoadPermission (src string) (error, *Permission) {
	p := new(Permission)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		return err, nil
	}
	if _, ok := r.permissions[p.Name]; ok {
		return nil, nil
	}
	r.permissions[p.Name] = p
	return nil, p
}