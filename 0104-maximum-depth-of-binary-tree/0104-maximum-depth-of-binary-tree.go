/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := 0
    right := 0

    if root.Left != nil {
        left = maxDepth(root.Left)
    }
    if root.Right != nil {
        right = maxDepth(root.Right)
    }
    return 1 + max(left, right)
}