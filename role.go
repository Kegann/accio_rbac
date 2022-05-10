package src

type Role struct {
	name        string
	//permissions []*Permission
	permissionMap map[string]*bLinkedNode
	//head, tail *bLinkedNode
}

// create a new role
func (r *Rbac) NewRole(name string, permissions []*Permission) *Role {
	new_role := &Role{
		name: name,
		// permissions: permissions,
		permissionMap: map[string]*bLinkedNode{}}
		//head: initbLinkNode(&Permission{}),
		//tail: initbLinkNode(&Permission{})}
	//new_role.head.next = new_role.tail
	//new_role.tail.next = new_role.head
	//r.roles = append(r.roles, new_role)
	for _, permission := range permissions {
		new_role.assign(permission)
	}
	return new_role
}

// check if role has a permission
func (role *Role) permit(comp_permission *Permission) bool {
	if _, ok := role.permissionMap[comp_permission.name]; ok {
		return true
	}
	return false
}

// add a permission for a role
func (role *Role) assign(new_permission *Permission) {
	if _, ok := role.permissionMap[new_permission.name]; !ok {
		permissionNode := initbLinkNode(new_permission)
		role.permissionMap[new_permission.name] = permissionNode
		//role.add2Head(permissionNode)
	} else {
		return
	}
	return
}

// remove role's permission
func (role *Role) revoke(revoke_permission *Permission) {
	if _, ok := role.permissionMap[revoke_permission.name]; !ok {
		return
	}
	//node := role.permissionMap[revoke_permission.name]
	//role.removeNode(node)
	delete(role.permissionMap, revoke_permission.name)
	return
}
