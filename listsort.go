package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	return insertionSort(l)
}

func findMax(l *NodeI) *NodeI {
	max := l
	index_i := l

	for index_i != nil {
		if index_i.Data > max.Data {
			max = index_i
		}
		index_i = index_i.Next
	}
	return max
}

func NodeRemove(l *NodeI, data int) {
	var prev *NodeI
	var next *NodeI

	current := l
	prev = nil
	next = nil

	for current != nil {
		next = current.Next
		if current.Data == data {
			if prev == nil {
				// Im in the head
				l = l.Next
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
func NodePushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
func insertionSort(head *NodeI) *NodeI {
	if head == nil || head.Next == nil {
		return head
	}

	sortedHead := head
	current := head.Next
	sortedHead.Next = nil

	for current != nil {
		next := current.Next
		sortedHead = insertNode(sortedHead, current)
		current = next
	}

	return sortedHead
}

func insertNode(sortedHead, newNode *NodeI) *NodeI {
	if newNode.Data < sortedHead.Data {
		newNode.Next = sortedHead
		return newNode
	}

	curr := sortedHead
	for curr.Next != nil && curr.Next.Data < newNode.Data {
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode

	return sortedHead
}
