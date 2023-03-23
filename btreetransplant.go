package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == node && root.Left != nil {
		rplc.Parent = root
		root.Left = rplc
	} else if root.Right == node && root.Right != nil {
		rplc.Parent = root
		root.Right = rplc
	} else {
		BTreeTransplant(root.Left, node, rplc)
		BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
