package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if root.Right != nil {
			Min := BTreeMin(root.Right)
			root.Data = Min.Data
			node = Min
		} else if root.Left != nil {
			Max := BTreeMin(root.Left)
			root.Data = Max.Data
			node = Max
		} else {
			return nil
		}
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else {
		root.Right = BTreeDeleteNode(root.Right, node)
	}

	return root
}

func BTremoveMax(root *TreeNode) {
	if root.Right == nil {
		root = nil
		return
	}
	BTremoveMax(root.Right)
}
func BTremoveMin(root *TreeNode) {
	if root.Left == nil {
		root = nil
		return
	}
	BTremoveMin(root.Left)
}
