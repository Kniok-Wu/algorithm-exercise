package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := 0
	for idx, val := range nums {
		if val > nums[max] {
			max = idx
		}
	}

	root := &TreeNode{
		Val:   nums[max],
		Left:  constructMaximumBinaryTree(nums[:max]),
		Right: constructMaximumBinaryTree(nums[max+1:]),
	}

	return root
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	res := constructMaximumBinaryTree(nums)
	fmt.Println(res)
}
