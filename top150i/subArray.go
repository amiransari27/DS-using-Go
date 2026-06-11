package top150i

/*
Write a function to find if B is a sub-array of A.
For B to be called a sub-array of A, the elements of B shall occur in A in consecutive indices
(but they need NOT be in the same order).
A = {3,2,7,1,4,6} and B = {7,1,4}. Here, B is a sub-array of A  ==> TRUE
A = {3,2,7,1,4,6} and B = {4,7,1}. Here also, B is a sub array of A  ==> TRUE
A = {3,2,7,1,4,6} and B = {3,1,4}. Here, B is not a sub array of A  ==> FALSE
A = {3,2,7,7,4,6} and B = {7,4,4}. Here, B is not a sub array of A  ==> FALSE
A = {2,2,2,2,2,7} and B = {2,2,2,2,7}. Here, B is a sub array of A  ==> TRUE
*/

func isSubArray(arrA []int, arrB []int) bool {
	m := len(arrA)
	n := len(arrB)

	i := 0
	j := i + n
	for j <= m {
		if checkB(arrA, arrB, i, j) {
			return true
		}
		i++
		j = i + n
	}
	return false
}

func checkB(arrA []int, arrB []int, l int, r int) bool {

	bMap := make(map[int]int)

	for _, v := range arrB {
		bMap[v]++
	}

	for i := l; i < r; i++ {
		bMap[arrA[i]]--
	}

	for _, v := range bMap {
		if v < 0 {
			return false
		}
	}

	return true
}
