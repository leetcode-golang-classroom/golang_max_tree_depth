package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return FindMaxDepth(root)
}

func FindMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(FindMaxDepth(root.Left), FindMaxDepth(root.Right)) + 1
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
