package main

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	i, j := 0, len(s)-1

	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func partition(s string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	n := len(s)

	var backTracking func(int)

	backTracking = func(i int) {
		if i == n {
			result = append(result, append([]string{}, path...))
			return
		}

		for j := i; j < n; j++ {
			if isPalindrome(s[i : j+1]) {
				path = append(path, s[i:j+1])
				backTracking(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backTracking(0)

	return result
}

func main() {
	partition("abbab")
}
