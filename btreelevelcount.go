package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left_height := BTreeLevelCount(root.Left)
	right_height := BTreeLevelCount(root.Right)

	return max(left_height, right_height) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
