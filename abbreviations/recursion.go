package abbreviations

import "fmt"

func solution(str string, str_sf string, count int, pos int) {

	if pos == len(str) {
		if count == 0 {
			fmt.Println(str_sf)
		} else {
			fmt.Printf("%s%d\n", str_sf, count)
		}
		return
	}

	if count > 0 {
		solution(str, fmt.Sprintf("%s%d%c", str_sf, count, str[pos]), 0, pos+1)
	} else {
		solution(str, fmt.Sprintf("%s%c", str_sf, str[pos]), 0, pos+1)
	}
	solution(str, str_sf, count+1, pos+1)
}

func ExeRecursion() {
	str := "pep"

	solution(str, "", 0, 0)
}
