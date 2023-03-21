package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	var next *NodeL

	current := l.Head
	prev = nil
	next = nil

	for current != nil {
		next = current.Next
		if current.Data == data_ref {
			if prev == nil {
				// Im in the head
				next = l.Head.Next
				//l.Head = nil
				l.Head = next
			} else {
				prev.Next = next

			}
		}
		prev = current
		current = next
	}
}
