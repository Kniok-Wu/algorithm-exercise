package main

import "fmt"

func BackTracking(target int, left []int) [][]int {
	result := make([][]int, 0)

	if !(target <= 0 || len(left) == 0) {

		for idx, c := range left {
			if target%c == 0 { // 可以取 n 个 m
				resultItem := make([]int, target/c)
				for i := 0; i < target/c; i += 1 {
					resultItem[i] = c
				}

				result = append(result, resultItem)
			}

			if len(left)-1 > 0 {
				resultItem := []int{c}
				for target_ := target - c; target_ > 0; target_ -= c {
					for _, each := range BackTracking(target_, left[idx+1:]) {
						result = append(result, append((append([]int{}, resultItem...)), each...))
					}
					resultItem = append(resultItem, c)
				}
			}
		}
	}

	return result
}

func combinationSum(candidates []int, target int) [][]int {
	return BackTracking(target, candidates)
}

func main() {
	for _, each := range combinationSum([]int{2, 3, 5}, 8) {
		fmt.Println(each)
	}
}
