package main

import (
	"fmt"
)

//树节点
type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func bfs(root *TreeNode) int {
	//初始化队列
	que := make([]*TreeNode, 0)
	que = append(que, root)
	dep := 1 //深度

	//循环队列所有节点
	for len(que) != 0 {
		curSize := len(que)
		//循环遍历每层节点
		for i := 0; i < curSize; i++ {
			//取出队列首节点
			tnode := que[0]
			fmt.Println(tnode)
			que = que[1:]
			for /*当前节点所有方向*/ {
				//对移动后的节点作坐标值变化

				//无法到达点，是否为边界，障碍物

				//终点判断，题目所求的点

				//入队列
			}
		}
		//深度变化（层数或步数）
		dep++
	}
	return -1
}
func main() {
	fmt.Println()
}
