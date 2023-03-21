package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	curr := l
	for curr.Next != nil && data_ref > curr.Next.Data {
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode

	return l
}
