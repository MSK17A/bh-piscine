package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListLast(l *List) interface{} {
	return l.Tail.Data
}
