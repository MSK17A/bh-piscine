package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	// Change lanes? GO TO THE RIGHT NODE
	BTreeApplyInorder(root.Right, f)
}
