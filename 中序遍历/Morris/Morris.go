package main

type TreeNode struct {
	val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var predecessor *TreeNode
	for root != nil {
		if root.LeftNode == nil { //如果x无左孩子
			res = res.append(res, root.val) //则将x的值加入数组
			root = root.RightNode           //再访问x的右孩子
		} else {
			//x有左孩子
			predecessor = root.LeftNode
			for predecessor.RightNode != root && predecessor.RightNode != nil {
				predecessor = predecessor.RightNode
			} //找到x的左子树中最右的节点
			if predecessor.RightNode == nil {
				predecessor.RightNode = root //如果pre右孩子为空，则将pre的右孩子指向x
				root = root.LeftNode         //然后访问x的左孩子
			}
			if predecessor.RightNode == root { //如果pre右孩子不为空（指向x）
				res = res.append(res, root) //则x输出
				predecessor.RightNode = nil //pre的右孩子置空
				root = root.RightNode       //再访问x的右孩子
			}
		}
	}
	return
}
