package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode = nil
var record = make([][]int, 0)

func compare() {
	if l := len(record); l > 1 {
		if record[l-1][1] > record[l-2][1] { // 删除前方元素
			record = [][]int{record[l-1]}
		} else if record[l-1][1] < record[l-2][1] { // 删除最后一个元素
			record = record[:l-1]
		}
	}
}

func find(root *TreeNode) {
	if root == nil {
		return
	}

	find(root.Left)

	if prev != nil { // 如果有前置节点
		if l := len(record); prev.Val == root.Val { // 如果当前元素与前一元素的值一致
			record[l-1][1] += 1
		} else { // 与之前的元素不一致
			compare()
			record = append(record, []int{root.Val, 1})
		}
	} else {
		record = append(record, []int{root.Val, 1})
	}

	prev = root
	find(root.Right)
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	find(root)

	compare()

	res := make([]int, len(record))
	for idx, val := range record {
		res[idx] = val[0]
	}
	prev = nil
	record = make([][]int, 0)
	return res
}

func main() {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	fmt.Println(findMode(root))
}
