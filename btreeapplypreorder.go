package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data) // Preorder apply
	BTreeApplyPreorder(root.Left, f)
	// Change lanes? GO TO THE RIGHT NODE
	BTreeApplyPreorder(root.Right, f)
}
