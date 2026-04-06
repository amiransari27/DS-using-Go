package top150i

func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)

	for i := 0; i < n; i++ {

		if gas[i] < cost[i] {
			continue
		}

		j := (i + 1) % n
		reqCap := cost[i]
		gasCap := gas[i]
		earnGas := gas[j]

		currGas := gasCap - reqCap + earnGas

		for j != i {

			if currGas < cost[j] {
				break
			}

			reqCap = cost[j]

			j = (j + 1) % n

			earnGas = gas[j]

			currGas = currGas - reqCap + earnGas

		}

		if i == j {
			return i
		}
	}

	return -1
}
