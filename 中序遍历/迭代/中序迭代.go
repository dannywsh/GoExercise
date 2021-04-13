package main

type TreeNode struct {
	val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	treeStack := make([]*TreeNode, 0)
	for root != nil || len(treeStack) > 0 { //如果节点不为空或栈不为空，则持续进栈
		for root != nil {
			treeStack = append(treeStack, root)
			root = root.LeftNode //遍历左节点
		}
		root = treeStack[len(treeStack)-1]       //左遍历到头惹，则栈顶为最后一个元素
		treeStack = treeStack[:len(treeStack)-1] //出栈
		res = append(res, root.val)              //存取值
		root = root.RightNode                    //开始遍历右节点惹
	}
	return
}
