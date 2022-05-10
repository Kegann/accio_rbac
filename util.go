package src

type bLinkedNode struct {
	permission *Permission
	prev, next *bLinkedNode
}

func initbLinkNode(val *Permission) *bLinkedNode {
	return &bLinkedNode{
		permission: val}
}

func (role *Role) removeNode(node *bLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(role.permissionMap, node.permission.name)
}

//func (role *Role) add2Head(node *bLinkedNode) {
//	node.prev = role.head
//	node.next = role.head.next
//	role.head.next.prev = node
//	role.head.next = node
//}