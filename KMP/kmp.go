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

	i, j := 0, 0
	for i < len(str) {
		if str[i] == subStr[j] {
			i += 1
			j += 1
		} else if j == 0 {
			i += 1
		} else {
			j = prefix[j-1]
		}

		if j == len(subStr) {
			return i - j
		}
	}

	return -1
}
