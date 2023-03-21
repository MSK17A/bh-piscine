package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListReverse(l *List) {
	var curr *NodeL
	var next *NodeL
	var prev *NodeL

	// Initial step
	curr = l.Head
	prev = nil
	l.Tail = curr
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
