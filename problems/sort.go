package problems

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		right := len(arr) - 1
		left := right - 1

		// for left > -1 {
		for (left - i) > -1 {
			if arr[left] > arr[right] {
				tmp := arr[left]
				arr[left] = arr[right]
				arr[right] = tmp
			}
			right--
			left--
		}

	}

	return arr
}

func BubbleSort2(arr []int) []int {

	for swapped := true; swapped; { // this is how we can achieve do while functionality in golang
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				swapped = true
			}
		}
	}

	return arr
}

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		small := i

		for j := i; j < len(arr); j++ {
			if arr[small] > arr[j] {
				small = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[small]
		arr[small] = tmp
	}

	return arr
}

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[i] < arr[i-1] {
				tmp := arr[i-1]
				arr[i-1] = arr[i]
				arr[i] = tmp
				j--
			}
			break
		}
	}

	return arr
}
