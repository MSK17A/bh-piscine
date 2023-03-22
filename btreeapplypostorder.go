package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
	// Change lanes? GO TO THE RIGHT NODE
	BTreeApplyPostorder(root.Right, f)
	f(root.Data) // Postorder apply
}
