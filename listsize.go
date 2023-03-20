package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListSize(l *List) int {
	size := 0
	tempHead := l.Head

	for tempHead != nil {
		size++
		tempHead = tempHead.Next
	}
	return size
}
