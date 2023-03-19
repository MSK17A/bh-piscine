package piscine

// Commented because already defined in listpushback.go (in package piscine)
/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListPushFront(l *List, data interface{}) {
	tempNode := NodeL{Data: data, Next: l.Head}

	l.Head = &tempNode
}
