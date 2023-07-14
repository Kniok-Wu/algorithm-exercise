package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := make([]string, 0)

	res = append(res, binaryTreePaths(root.Left)...)

	res = append(res, binaryTreePaths(root.Right)...)

	if len(res) == 0 {
		res = append(res, strconv.Itoa(root.Val))
	} else {
		for i := 0; i < len(res); i++ {
			res[i] = strconv.Itoa(root.Val) + "->" + res[i]
		}
	}

	return res
}
