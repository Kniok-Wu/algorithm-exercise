package KMP

func KMP(str, subStr string) int {
	prefix := make([]int, len(subStr)) // 前缀表
	for i := 0; i < len(subStr); i += 1 {
		k := i
		for k > 0 {
			if subStr[:k] == subStr[i-k+1:i+1] {
				break
			}
			k--
		}
		prefix[i] = k
	}

	for i := 0; i < len(str); {
		j := 0
		for ; j < len(subStr); j++ {
			if subStr[j] != str[i+j] {
				if j == 0 {
					i += 1
				} else {
					i += prefix[j-1] + 1
				}
				break
			}
		}
		if j == len(subStr) {
			return i
		}
	}

	return -1
}
