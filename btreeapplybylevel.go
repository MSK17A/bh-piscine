package piscine

type QueueNode struct {
	treeNode TreeNode
	Next     *QueueNode
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

func QueuePushBack(l *Queue, treeNode *TreeNode) {
	tempNode := &QueueNode{treeNode: *treeNode}
	if l.Head == nil {
		l.Head = tempNode
		l.Tail = tempNode
		return
	}

	l.Tail.Next = tempNode
	l.Tail = l.Tail.Next
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	queue := &Queue{}
	QueuePushBack(queue, root)

	for i := queue.Head; i != nil; i = i.Next {
		f(i.treeNode.Data)
		if i.treeNode.Left != nil {
			QueuePushBack(queue, i.treeNode.Left)
		}
		if i.treeNode.Right != nil {
			QueuePushBack(queue, i.treeNode.Right)
		}
	}
}
