package dynamicProg

/*
Partition problem is to determine whether a given set can be partitioned into two subsets such that the sum of elements in both subsets is same.
Examples:

arr[] = {1, 5, 11, 5}
Output: true
The array can be partitioned as {1, 5, 5} and {11}
*/

func EqualSumPartition(arr []int) bool {
	totalSum := 0
	for _, num := range arr {
		totalSum += num
	}

	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	return SubsetSum(arr, target)
}
