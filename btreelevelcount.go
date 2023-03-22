package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	BTreeLevelCount(root.Left)
	return BTreeLevelCount(root.Right) + 1
}
