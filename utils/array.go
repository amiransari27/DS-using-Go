package utils

func LowerBound(arr []int, num int) int {

	left := 0
	right := len(arr)

	for left < right {

		mid := left + (right-left)/2

		if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return left

}

func UpperBound(arr []int, target int) int {
	left := 0
	right := len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
