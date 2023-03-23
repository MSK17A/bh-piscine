package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isValidBST(root, nil, nil)
}

func isValidBST(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}
	return isValidBST(node.Left, min, node) && isValidBST(node.Right, node, max)
}
