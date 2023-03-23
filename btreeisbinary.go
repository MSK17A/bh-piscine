package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if BTreeIsBinary(root.Left) == false || BTreeIsBinary(root.Right) == false {
		return false
	}
	return true
}

/*func isValidBST(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}
	return isValidBST(node.Left, min, node) && isValidBST(node.Right, node, max)
}*/
