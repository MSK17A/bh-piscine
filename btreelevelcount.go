package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	BTreeLevelCount(root.Left)
	return 1 + BTreeLevelCount(root.Right)
}
