package problems

import "fmt"

func Cartisian(set1 []int, set2 []int) [][]int {
	var product [][]int
	for i := 0; i < len(set1); i++ {
		for j := 0; j < len(set2); j++ {
			product = append(product, []int{set1[i], set2[j]})
		}
	}
	return product
}

// go generics

func Min[T int | float64](x T, y T) T {
	if x > y {
		return y
	}
	return x
}

type minType interface {
	float64 | int
}

func Min2[T minType](x T, y T) T {
	if x > y {
		return y
	}
	return x
}

// func GetMe[T any](x T, y T) T {
func GetMe[T interface{}](x T, y T) T {
	return x
}

// func test[T float64](x T) T { // this case will throw error 
func test[T ~float64](x T) T {
	return x
}

func Test() {
	type superFloat float64
	fmt.Println(test[superFloat](41.1))
}


// copied from main 
// fmt.Println(problems.Cartisian([]int{1, 2}, []int{2, 3, 4}))
// 	fmt.Println(problems.Min(6.2, 5))
// 	fmt.Println(problems.Min2(4.2, 5.1))
// 	fmt.Println(problems.Min2(25, 17))
// 	fmt.Println(problems.GetMe[string]("amir", "ansari"))
// 	fmt.Println(problems.GetMe(1, 2))
// 	problems.Test()
