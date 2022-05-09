# golang_max_tree_depth

Given the `root` of a binary tree, return *its maximum depth*.

A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 3

```

**Example 2:**

```
Input: root = [1,null,2]
Output: 2

```

**Constraints:**

- The number of nodes in the tree is in the range `[$0, 10^4$]`.
- `100 <= Node.val <= 100`

## 解析

給一個二元樹的根結點 root 

要找出二元樹的最大深度

每個結點都是由左結點跟右結點所組成

所以對於每個結點其最大深度 = Max(左結點最大深度，右結點最大深度) + 1(因為加上目前這層深度)

當該結點是空值時，最大深度定義是 0(因為不存在)

## 程式碼
```go
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
```
## 困難點

1. 要思考出二元樹最大深度的關係式
2. 知道邊界條件

## Solve Point

- [x]  Understand What problem need to solve
- [x]  Analysis Complexity