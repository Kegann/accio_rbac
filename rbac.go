package src

type Rbac struct {
	users map[int64]*User
	permissions map[string]*Permission
	roles map[string]*Role
}

type Rbaci interface {
	NewPermission(name string) *Permission
	NewRole(name string, permissions []*Permission) *Role
	NewUser(id int64, name string, roles []*Role) *User
}

// persist data into db; support mysql, pgsql
// ptype: 1: mysql; 2: pgsql
func (*Rbac) persist2Db(ptype int) bool {
	if (ptype != 1 && ptype != 2) {
		return false
	}

	return false
}

// 从json导入
func (*Rbac) loadsFromJson(data string) bool {
	return false
}


// 从db导入
func (*Rbac) loadsFromDb(db interface{}) bool {
	return false
}
