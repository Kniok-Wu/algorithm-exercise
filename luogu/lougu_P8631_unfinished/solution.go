package main

import (
	"fmt"
	"strings"
)

func ManacherStr(str string) (mStr string) {
	var builder strings.Builder

	for _, b := range []byte(str) {
		builder.WriteString("#")
		builder.Write([]byte{b})
	}
	builder.WriteString("#")
	mStr = builder.String()
	return
}

func Manacher(str string, reverse bool) (count int) {
	d := make([]int, len(str)) // 初始化回文半径
	d[0] = 1                   // 设置首元素的回文半径为1

	record := make(map[string]int)

	l := 0
	r := 0 // 盒子半径

	for i := 1; i < len(str); i++ {
		mid := (l + r) / 2   // 盒子中点
		if i >= l && i < r { // 如果在盒子范围内
			d[i] = d[mid-(i-mid)] // 对应点位
		} else { // 不在盒子范围内 进行暴力搜索
			d[i] = 1
			for j := 1; i-j >= 0 && i+j < len(str); j++ {
				if str[i-j] == str[i+j] {
					d[i] += 1
				} else {
					break
				}
			} // 计算出当前节点的回文半径
			if d[i] > d[mid] { // 如果出现更长的回文子串 则更新盒子范围
				l = i - (d[i] - 1)
				r = i + (d[i] - 1)
			}
		}
	}

	for i := 0; i < len(d); i += 1 {
		v := d[i]
		key := strings.Replace(string(str[i-(v-1):i+v]), "#", "", -1)
		if key != "" {
			_, ok := record[key]
			if !ok {
				if len(key)%2 != 0 && !reverse || len(key)%2 == 0 && reverse {
					count += 1
				}
			}
		}
	}

	return
}

func main() {
	var max int
	var t int
	var str string

	fmt.Scan(&t)
	fmt.Scan(&str)

	for i := 1; i < len(str); i++ {
		subStr1 := string(str[:i])
		subStr2 := string(str[i:])

		//r1 := Manacher(ManacherStr(subStr1), false)
		//r2 := Manacher(ManacherStr(subStr2), true)
		//
		//fmt.Println(r1, r2)

		if result := Manacher(ManacherStr(subStr1), false) * Manacher(ManacherStr(subStr2), true); result > max {
			max = result
		}
	}
	fmt.Println(max)
}
