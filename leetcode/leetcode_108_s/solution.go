package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	if len(nums) == 0 {
		return nil
	}

	var mid int = len(nums) / 2

	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func main() {
	a := []int{1, 2, 3}

	fmt.Println(len(a[3:]))
}
