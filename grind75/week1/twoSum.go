package week1

import (
	"fmt"
	"sort"
)

func TwoSum(nums []int, target int) []int {
	res := make([]int, 0)
	sort.Ints(nums)

	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			fmt.Println("Found it:", nums[left], nums[right], left, right)
			res = append(res, left, right)
			break
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}
	return res
}

func TwoSum1(nums []int, target int) []int {
	res := make([]int, 0)

	m := make(map[int]int, 0)

	for i, num := range nums {
		complement := target - num
		if idx, ok := m[complement]; ok {
			res = append(res, idx, i)
			return res
		}
		m[num] = i
	}

	return nil
}
