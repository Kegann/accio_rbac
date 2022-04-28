package src

type Permission struct{
	name string
}

func NewPermission(name string) *Permission {
	return &Permission{name: name}
}