/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        l := dfs(node.Left)
        if l == -1 {
            return -1
        }
        r := dfs(node.Right)
        if r == -1 || abs(l-r) > 1 {
            return -1
        }
        return max(l, r) + 1
    }
    return dfs(root) != -1
}

func abs(x int) int { if x < 0 { return -x }; return x }
func max(a, b int) int { if a > b { return a }; return b }