package praxis

//#翻转一棵二叉树。

//示例：

//输入：

//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//输出：

//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func invertTree(root *TreeNode) *TreeNode {
	return overturn(root)
}

func overturn(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	node.Left, node.Right = node.Right, node.Left
	overturn(node.Left)
	overturn(node.Right)
	return node
}
