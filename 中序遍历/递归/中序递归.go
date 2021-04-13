package main

type TreeNode struct {
	val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.LeftNode)
		res = append(res, node.val)
		inorder(node.RightNode)
	}
	inorder(root)
	return
}
