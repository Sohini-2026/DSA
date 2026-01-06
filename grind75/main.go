package main

import (
	"fmt"

	"github.com/Sohini-2026/DSA/grind75/week1"
)

// func Marea(heights []int) int {
// 	// Your code goes here
// 	left, right := 0, len(heights)-1
// 	currmaxArea := 0
// 	maxArea := 0

// 	for left < right {
// 		fmt.Println(left, right, heights[left], heights[right], maxArea)
// 		if heights[right] > heights[left] {
// 			currmaxArea = (right - left) * heights[left]
// 			left++
// 		} else {
// 			currmaxArea = (right - left) * heights[right]
// 			right--
// 		}
// 		if currmaxArea > maxArea {
// 			maxArea = currmaxArea
// 		}
// 	}

// 	return maxArea
// }

// func main() {
// 	fmt.Println(Marea([]int{3, 4, 1, 2, 2, 4, 1, 3, 2}))
// }

func main() {
	res := week1.TwoSum1([]int{3, 2, 4}, 6)
	fmt.Printf("Result:::%+v", res)

	fmt.Printf("ValidParentheses1:::%+v", week1.ValidParentheses1("()[]{}"))
	fmt.Printf("ValidParentheses1:::%+v", week1.ValidParentheses1("[{()}]"))
	week1.PrintLL(week1.MergeTwoLists(&week1.ListNode{1, &week1.ListNode{2, &week1.ListNode{4, nil}}}, &week1.ListNode{1, &week1.ListNode{3, &week1.ListNode{4, nil}}}))
	//fmt.Printf("MergeTwoLists:::%+v", mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
	fmt.Println("Max profit::::", week1.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("isPalindrome::::", week1.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("%+v", week1.PrintTree(
		week1.InvertBinaryTree(
			&week1.TreeNode{
				Val: 1,
				Left: &week1.TreeNode{
					Val:   2,
					Left:  &week1.TreeNode{Val: 4},
					Right: &week1.TreeNode{Val: 7},
				},
				Right: &week1.TreeNode{
					Val:   3,
					Left:  &week1.TreeNode{Val: 5},
					Right: &week1.TreeNode{Val: 6},
				},
			},
		),
	))
	fmt.Println(week1.IsAnagram("anagram", "nagarami"))
	fmt.Printf("Binary Search::::%d", week1.BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Printf("%+v", week1.LowestcommonAncestor2(&week1.TreeNode{
		Val: 6,
		Left: &week1.TreeNode{
			Val:   2,
			Left:  &week1.TreeNode{Val: 0},
			Right: &week1.TreeNode{Val: 4},
		},
		Right: &week1.TreeNode{
			Val:   8,
			Left:  &week1.TreeNode{Val: 7},
			Right: &week1.TreeNode{Val: 9},
		},
	}, &week1.TreeNode{Val: 2}, &week1.TreeNode{Val: 4}))
}
