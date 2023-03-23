package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left == node {
		root.Left = rplc
		root.Left.Parent = root

	} else if root.Right != nil && root.Right == node {
		root.Right = rplc
		root.Right.Parent = root
	} else {
		BTreeTransplant(root.Left, node, rplc)
		BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
