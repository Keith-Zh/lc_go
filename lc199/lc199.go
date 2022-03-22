/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	//edge case
	if root == nil {
		return res
	}
	//这道题的逻辑就是进行bfs，每次记录最后一位
	queue := []*TreeNode{root}
	var levelEnd int
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curNode := queue[i]
			//更新level的最后一个node
			levelEnd = curNode.Val
			//add left and right child
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		//更新queue，把所有pop掉的node都从queue里面移除
		queue = queue[levelSize:]
		//走完这一层把levelEnd加入到res里面去
		res = append(res, levelEnd)
	}
	return res
}