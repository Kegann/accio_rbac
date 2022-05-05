package src

type Role struct {
	name        string
	permissions []*Permission

	permissionMap map[string]*bLinkedNode
	head, tail *bLinkedNode
}

func (r *Rbac) NewRole(name string, permissions []*Permission) *Role {
	new_role := &Role{
		name: name,
		permissions: permissions,
		permissionMap: map[string]*bLinkedNode{},
		head: initbLinkNode(&Permission{}),
		tail: initbLinkNode(&Permission{})}
	new_role.head.next = new_role.tail
	new_role.tail.next = new_role.head
	//r.roles = append(r.roles, new_role)
	for _, permission := range permissions {
		new_role.assign(permission)
	}
	return new_role
}

func (role *Role) permit(comp_permission *Permission) bool {
	//for _, permission := range role.permissions {
	//	if comp_permission.name == permission.name {
	//		return true
	//	}
	//}
	if _, ok := role.permissionMap[comp_permission.name]; ok {
		return true
	}
	return false
}

func (role *Role) assign(new_permission *Permission) {
	// check if exist
	// todo optimize hash+link?
	//for _, permission := range role.permissions {
	//	if new_permission.name == permission.name {
	//		return
	//	}
	//}
	//role.permissions = append(role.permissions, new_permission)

	if _, ok := role.permissionMap[new_permission.name]; !ok {
		permissionNode := initbLinkNode(new_permission)
		role.permissionMap[new_permission.name] = permissionNode
		role.add2Head(permissionNode)
	} else {
		return
	}
	return
}

// remove role's permission
func (role *Role) revoke(revoke_permission *Permission) {
	// check if exist before delete
	// todo optimize hash+link?
	//for i := 0; i < len(role.permissions); i++ {
	//	if revoke_permission.name == role.permissions[i].name {
	//		role.permissions = append(role.permissions[:i], role.permissions[i+1:]...)
	//		return
	//	}
	//}

	if _, ok := role.permissionMap[revoke_permission.name]; !ok {
		return
	}
	node := role.permissionMap[revoke_permission.name]
	role.removeNode(node)
	return
}
