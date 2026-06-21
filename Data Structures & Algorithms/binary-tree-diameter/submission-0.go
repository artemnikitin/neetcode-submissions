/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, diameter)
	right := dfs(root.Right, diameter)
	*diameter = max(*diameter, left+right)
	return max(left, right) + 1
}
