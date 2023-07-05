package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)

	for i := 0; i < len(s); i++ {
		arr[int(s[i])-97] += 1
		arr[int(t[i])-97] -= 1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("abc", "bad"))
}
