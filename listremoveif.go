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
				l.Head = l.Head.Next
				current = nil
			} else {
				prev.Next = next
				current = nil
			}
		} else {
			// Inside an else statement to make sure the prev won't hold the address of the current deleted node
			prev = current
		}
		current = next
	}
}
