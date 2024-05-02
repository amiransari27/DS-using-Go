package top150

func CanCompleteCircuit(gas []int, cost []int) int {
	totalDiff, fuel, index := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalDiff += diff
		fuel += diff
		if fuel < 0 {
			index = i + 1
			fuel = 0
		}
	}

	if totalDiff < 0 {
		return -1
	} else {
		return index
	}

}
