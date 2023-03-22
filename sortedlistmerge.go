package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var head, tail *NodeI
	if n1 == nil && n2 == nil {
		return nil
	}

	if n2 == nil {
		return n1
	}

	if n1 == nil {
		return n2
	}

	// set the head and tail to the smaller node
	if n1.Data < n2.Data {
		head, tail = n1, n1
		n1 = n1.Next
	} else {
		head, tail = n2, n2
		n2 = n2.Next
	}

	// merge the two lists
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			tail.Next = n1
			n1 = n1.Next
		} else {
			tail.Next = n2
			n2 = n2.Next
		}
		tail = tail.Next
	}

	// append the remaining nodes
	if n1 != nil {
		tail.Next = n1
	} else if n2 != nil {
		tail.Next = n2
	}

	return head
}
