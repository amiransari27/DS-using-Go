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
			if arr[j] < arr[j-1] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
				j--
			} else {
				break
			}
		}
	}

	return arr
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)-1]
	leftArr := []int{}
	rightArr := []int{}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > pivot {
			rightArr = append(rightArr, arr[i])
		} else {
			leftArr = append(leftArr, arr[i])
		}
	}

	left := QuickSort(leftArr)
	right := QuickSort(rightArr)

	res := []int{}
	res = append(res, left...)
	res = append(res, pivot)
	res = append(res, right...)

	return res
}

func QuickSort2(arr []int) []int {

	quickSort(arr, 0, len(arr)-1)

	return arr
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partistion(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partistion(arr []int, low int, high int) int {

	pivot := arr[high]
	position := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			position++
			tmp := arr[position]
			arr[position] = arr[j]
			arr[j] = tmp
		}
	}
	position++
	tmp := arr[position]
	arr[position] = pivot
	arr[high] = tmp
	return position
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := append([]int{}, arr[:mid]...)
	rigthArr := append([]int{}, arr[mid:]...)

	left := MergeSort(leftArr)
	right := MergeSort(rigthArr)
	resArr := []int{}
	for len(left) > 0 && len(right) > 0 {
		var tmp int
		if left[0] < right[0] {
			tmp = left[0]
			left = append([]int{}, left[1:]...)
		} else {
			tmp = right[0]
			right = append([]int{}, right[1:]...)
		}
		resArr = append(resArr, tmp)
	}

	resArr = append(resArr, left...)
	resArr = append(resArr, right...)

	return resArr
}
