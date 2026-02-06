/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func sumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(node *TreeNode, current int) int {
    if node == nil {
        return 0
    }
    
    current = current*10 + node.Val
    if node.Left == nil && node.Right == nil {
        return current 
    }

    leftSum  := dfs(node.Left,  current)
    rightSum := dfs(node.Right, current)

    return leftSum + rightSum

}