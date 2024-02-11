package leetcode75

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func PathSum(root *TreeNode, targetSum int) int {

	pathSumCount := 0
	memo := map[int]int{
		0: 1,
	}

	// pathSumSol(root, targetSum, &pathSumCount)
	findSumOptimized(root, targetSum, 0, &pathSumCount, memo)
	return pathSumCount
}

func findSumOptimized(root *TreeNode, target int, currentSum int, pathSumCount *int, memo map[int]int) {
	if root == nil {
		return
	}

	currentSum += root.Val
	diff := currentSum - target

	*pathSumCount += memo[diff]
	memo[currentSum]++
	findSumOptimized(root.Left, target, currentSum, pathSumCount, memo)
	findSumOptimized(root.Right, target, currentSum, pathSumCount, memo)
	memo[currentSum]--
}

func pathSumSol(root *TreeNode, targetSum int, pathSumCount *int) {
	if root == nil {
		return
	}

	findSum(root, targetSum, 0, pathSumCount)
	pathSumSol(root.Left, targetSum, pathSumCount)
	pathSumSol(root.Right, targetSum, pathSumCount)
}

func findSum(root *TreeNode, targetSum int, currentSum int, pathSumCount *int) {
	if root == nil {
		return
	}
	currentSum += root.Val
	if currentSum == targetSum {
		*pathSumCount++
	}
	findSum(root.Left, targetSum, currentSum, pathSumCount)
	findSum(root.Right, targetSum, currentSum, pathSumCount)
}
