package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func NodeTail(head *NodeL) *NodeL {
	if head.Next == nil {
		return head
	}
	return NodeTail(head.Next)
}

func ListPushBack(l *List, data interface{}) {
	tempNode := NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = &tempNode
		return
	}

	l.Tail = NodeTail(l.Head)
	l.Tail.Next = &tempNode
}
