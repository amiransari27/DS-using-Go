package leetcode75

import (
	"fmt"
	"math"
)

func AsteroidCollision(asteroids []int) []int {

	stack := make([]int, 1)
	stack[0] = asteroids[0]

	for i := 1; i < len(asteroids); {

		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			if len(stack) == 0 {
				stack = append(stack, asteroids[i])
				i++
				continue
			}
			lastEle := stack[len(stack)-1]
			currentEle := asteroids[i]

			fmt.Println(math.Sqrt(float64(currentEle * currentEle)))
			fmt.Println(math.Sqrt(float64(lastEle * lastEle)))

			if lastEle == currentEle {
				stack = append(stack, asteroids[i])
			} else if lastEle < 0 && currentEle < 0 {
				stack = append(stack, asteroids[i])
			} else if lastEle > 0 && math.Sqrt(float64(currentEle*currentEle)) == math.Sqrt(float64(lastEle*lastEle)) {
				stack = stack[:len(stack)-1]
			} else if lastEle > 0 && math.Sqrt(float64(currentEle*currentEle)) > math.Sqrt(float64(lastEle*lastEle)) {
				stack = stack[:len(stack)-1]
				continue
			}
		}

		i++
	}

	return stack
}

func AsteroidCollisionOptimize(asteroids []int) []int {

	stack := make([]int, 0)

	for _, a := range asteroids {

		isAppend := true
		for len(stack) > 0 && stack[len(stack)-1] > 0 && a < 0 {
			sum := stack[len(stack)-1] + a

			if sum == 0 {
				stack = stack[:len(stack)-1]
				isAppend = false
				break
			} else if sum < 0 {
				stack = stack[:len(stack)-1]
			} else {
				isAppend = false
				break
			}
		}
		if isAppend {
			stack = append(stack, a)
		}

	}

	return stack
}
