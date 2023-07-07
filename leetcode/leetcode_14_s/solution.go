package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	i := 0
	for {
		for j := 0; j+1 < len(strs); j++ {
			if !(len(strs[j]) >= i && len(strs[j+1]) >= i) || strs[j+1][:i] != strs[j][:i] { // 必须均大于且前i部分相同
				i -= 1
				if i <= 0 {
					i = 0
				}
				goto RETURN
			}
		}
		i += 1
	}

RETURN:
	return strs[0][:i]
}

func longestCommonPrefix_2(strs []string) string { // 垃圾中的垃圾
	if len(strs) == 1 {
		return strs[0]
	}

	shortest := 0

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(strs[shortest]) {
			shortest = i
		}
	}

	for i := 1; i <= len(strs[shortest]); i++ {
		for j := 0; j < len(strs); j++ {
			fmt.Println(strs[j][:i], strs[shortest][:i])
			if strs[j][:i] != strs[shortest][:i] {
				i -= 1
				if i <= 0 {
					i = 0
				}
				return strs[shortest][:i]
			}
		}
	}

	return strs[shortest]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"", ""}))
	fmt.Println(longestCommonPrefix_2([]string{"cir", "car"}))
	fmt.Println(longestCommonPrefix_2([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}
