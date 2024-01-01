package problems

func LinearSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	var left int = 0
	var right int = len(arr) - 1

	for left <= right {
		var mid int = (left + right) / 2

		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1

}

func BinarySearchRec(arr []int, target int, left int, right int) int {
	if len(arr) == 0 {
		return -1
	}
	if left > right {
		return -1
	}
	var mid int = (left + right) / 2
	if target == arr[mid] {
		return mid
	} else if target > arr[mid] {
		return BinarySearchRec(arr, target, mid+1, right)
	} else {
		return BinarySearchRec(arr, target, left, mid-1)
	}

}
