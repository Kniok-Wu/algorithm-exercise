package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := strings.Builder{}
	for j := len(s) - 1; j >= 0; {
		for j >= 0 && s[j] == ' ' {
			j--
		} // 跳过所有空格 或者 全部搜索完成

		if j == -1 { // 如果没有字符 则结束
			break
		}

		i := j - 1                  // 继续向前搜索
		for i >= 0 && s[i] != ' ' { // 搜索到没有字符 或者 搜索到空格
			i--
		}
		res.WriteString(s[i+1 : j+1]) // i 可能搜索到没有字符(i = -1) 或者 i指向空格，因此 i + 1, 已知j指向字符，那么必须包含j，根据左闭右开，j + 1
		res.WriteByte(' ')            // 写入一个单词 跟一个空格
		j = i
	}

	result := res.String()
	return result[:len(result)-1] // 移除最后一个空格
}

func main() {
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("        hello"))
	fmt.Println(reverseWords("world        hello     "))
}
