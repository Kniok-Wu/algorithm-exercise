package main

import (
	"fmt"
	"strconv"
	"strings"
)

func judge(s string) bool {
	addrInt, _ := strconv.Atoi(s)
	return !(addrInt > 255 || (s[0] == '0' && len(s) != 1))
}

func restoreIpAddresses_1(s string) []string {
	result := make([]string, 0)
	parts := make([]string, 4)

	for i := 0; i < len(s) && i < 3; i++ {
		parts[0] = s[:i+1]
		if !judge(parts[0]) {
			break
		}
		for j := i + 1; j < len(s) && j-i <= 3; j++ {
			parts[1] = s[i+1 : j+1]
			if !judge(parts[1]) {
				break
			}
			for k := j + 1; k < len(s) && k-j <= 3; k++ {
				parts[2] = s[j+1 : k+1]
				if !judge(parts[2]) {
					break
				}
				for l := k + 1; l < len(s) && l-k <= 3; l++ {
					parts[3] = s[k+1 : l+1]
					if !judge(parts[3]) {
						break
					}
					if l+1 == len(s) {
						result = append(result, strings.Join(parts, "."))
					}
				}
			}
		}
	}

	return result
}

func restoreIpAddresses_2(s string) []string {
	result := make([]string, 0)
	records := make([]string, 0)

	var backTracking func(int)

	backTracking = func(start int) {
		if len(records) == 4 || start == len(s) {
			if len(records) == 4 && start == len(s) { // 如果没有完整应用字符串
				result = append(result, strings.Join(records, "."))
			}
			return
		}

		for i := 1; i <= 3 && start+i <= len(s); i++ {
			// 如果不满足小于255或者以0开头 直接退出
			if !judge(s[start : start+i]) {
				break
			}
			records = append(records, s[start:start+i])
			backTracking(start + i)
			records = records[:len(records)-1]
		}
	}

	backTracking(0)

	return result
}

func main() {
	result := restoreIpAddresses_2("101023")

	for _, each := range result {
		fmt.Println(each)
	}
}
