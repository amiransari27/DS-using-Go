package leetcode75

func MinCostClimbingStairs(cost []int) int {

	memo := make([]int, 1001)

	for i, _ := range memo {
		memo[i] = -1
	}

	costOf0 := solveMinCostClimbingStairs(cost, 0, memo)
	costOf1 := solveMinCostClimbingStairs(cost, 1, memo)

	return min(costOf0, costOf1)
}

func solveMinCostClimbingStairs(cost []int, i int, memo []int) int {

	if i >= len(cost) {
		return 0
	}

	if memo[i] != -1 {
		return memo[i]
	}

	costOf1Step := cost[i] + solveMinCostClimbingStairs(cost, i+1, memo)
	costOf2Step := cost[i] + solveMinCostClimbingStairs(cost, i+2, memo)

	memo[i] = min(costOf1Step, costOf2Step)
	return memo[i]
}

func MinCostClimbingStairsOptimized(cost []int) int {

	cost = append(cost, 0)

	for i := 2; i < len(cost); i++ {
		cost[i] = cost[i] + min(cost[i-1], cost[i-2])
	}

	return cost[len(cost)-1]
}
