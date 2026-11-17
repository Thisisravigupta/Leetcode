/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
	res := 0

	if root.Left != nil {
		if checkLeafNode(root.Left) {
			res += root.Left.Val
		} else {
			res += sumOfLeftLeaves(root.Left)
		}
	}
	res += sumOfLeftLeaves(root.Right)

	return res
}

func checkLeafNode(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}