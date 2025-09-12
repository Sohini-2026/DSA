package main

import (
	"fmt"
	"sort"

	"github.com/Sohini-2026/DSA/baseDS"
	"github.com/Sohini-2026/DSA/binarySearch"
	"github.com/Sohini-2026/DSA/dynamicProg"
	"github.com/Sohini-2026/DSA/graph"
	"github.com/Sohini-2026/DSA/recursive"
	slidingwindow "github.com/Sohini-2026/DSA/slidingWindow"
)

func main() {

	//recursion : start
	fmt.Println("What did u return ?", recursive.Print1ToN(5))
	recursive.Print1ToN(5)
	fmt.Println("Fibonacci of 5 is ", recursive.Fibonacci(5))
	fmt.Println("Height of BT is ", recursive.FindHeightOfBT(&recursive.Node{Data: 1, Left: &recursive.Node{Data: 2}, Right: &recursive.Node{Data: 3}}))
	fmt.Println("Sorting array using recursion")
	//arr := []int{5, 2, 9, 1, 5, 6}
	arr := []int{1, 5, 2, 0}
	recursive.SortArray(&arr)
	fmt.Println("Sorted array is ", arr)
	fmt.Println("Sorting stack using recursion")
	stack := &baseDS.Stack{Items: []int{1, 5, 2, 0}}
	recursive.SortStack(stack)
	fmt.Println("Sorted stack is ", stack.Items)
	fmt.Println("Deleting middle element of stack using recursion")
	deleteStack := &baseDS.Stack{Items: []int{1, 2, 3, 4, 5}}
	recursive.DeleteMiddleElementOfStack(deleteStack)
	fmt.Println("Stack after deleting middle element is ", deleteStack.Items)
	fmt.Println("Reversing stack using recursion")
	reverseStack := &baseDS.Stack{Items: []int{1, 2, 3, 4, 5}}
	recursive.ReverseStack(reverseStack)
	fmt.Println("Reversed stack is ", reverseStack.Items)

	fmt.Println("Tower of Hanoi with 3 disks")
	recursive.TowerOfHanoi(3, "A", "C", "B")
	fmt.Println("Kth symbol in grammar for n=4 and k=3 is ", recursive.GetKthSymbolInGrammer(4, 1))

	fmt.Printf("Print all subsets of array %v\n", []string{"a", "b", "c"})
	recursive.PrintAllSubsets1("abc", "")

	m := make(map[string]bool)
	result := []string{}
	recursive.PrintUniqueSubsetsLexicographially("abb", "", m, &result)
	sort.Strings(result)
	fmt.Println(result)

	fmt.Printf("Permutation With Spaces %v\n", "abc")
	recursive.PermutationWithSpaces("a", "bc")

	fmt.Printf("Permutation With Case Change %v\n", "abc")
	recursive.PermutationWithCaseChange(" ", "abc")
	fmt.Printf("Letter Case Permutation %v\n", "a1B2")
	recursive.LetterCasePermutation("", "a1B2")

	fmt.Printf("Balanced Parenthesis for n=3\n")
	recursive.CreateBalancedParanthesis("", "", 3, 3)
	fmt.Printf("Is Balanced Parenthesis for string %s\n", "({[]})")
	fmt.Println(recursive.IsBalancedParanthesis("({[]})"))

	fmt.Printf("Get N bit Binary Number for n=3\n")
	recursive.NBitBinaryNumber(5, "", 0, 0)

	fmt.Printf("Get last index in josephus problem for n=7 and k=3\n")
	fmt.Println(recursive.JosepheusProblem(7, 3))

	// recursion : end

	//binary search : start
	fmt.Printf("Binary Search for element 5 in array %v\n", []int{1, 2, 3, 4, 5, 6, 7})
	res := binarySearch.SearchNumberInSortedArray([]int{1, 2, 3, 4, 5, 6, 7}, 5)
	fmt.Println("Element found at index:", res)
	res = binarySearch.SearchNumberInReverseSortedArray([]int{7, 6, 5, 3, 2, 1}, 1)
	fmt.Println("Element found at index:", res)

	res = binarySearch.SearchNumberWithUnknownSorting([]int{1, 2, 3}, 1)
	fmt.Println("Element found at index:", res)

	count := binarySearch.FindCountOfOccurrences([]int{1, 2, 3, 4, 5, 5, 5, 6}, 5)
	fmt.Println("Count of occurrences of 5 is:", count)

	res = binarySearch.CountNoOfTimesArrayIsRotated([]int{6, 7, 1, 2, 3, 4, 5})
	fmt.Println("Count of times array is rotated is:", res)
	res = binarySearch.CountNoOfTimesArrayIsRotated([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("For non rotated array - Count of times array is rotated is:", res)

	//res = binarySearch.CountNoOfTimesArrayIsRotated([]int{7, 6, 5, 4, 3, 2, 1})
	//fmt.Println("Completely rotated - Count of times array is rotated is:", res)

	res = binarySearch.FindElementInRotatedSortedArray([]int{6, 7, 1, 2, 3, 4, 5}, 3)
	fmt.Println("Element 3 found at index:", res)

	res = binarySearch.FindElementInNearlySortedArray([]int{5, 10, 30, 20, 40}, 20)
	fmt.Println("Element 5 found at index:", res)

	res = binarySearch.FindFloorOfElementInSortedArray([]int{1, 2, 3, 4, 8, 10}, 5)
	fmt.Println("Floor index::", res)

	res = binarySearch.FindCeilOfElementInSortedArray([]int{1, 2, 3, 4, 8, 10}, 5)
	fmt.Println("Ceil index::", res)

	ress := binarySearch.FindNextAlphabeticalElement([]string{"a", "b", "c", "d"}, "c")
	fmt.Println("Next alphabetical element after 'c' is:", ress)

	res = binarySearch.FindPositionOfElementInInfiniteSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20}, 5)
	fmt.Println("Position of element 5 in infinite sorted array is:", res)

	res = binarySearch.FindIndexOf1InBSInfiniteArray([]int{0, 0, 0, 0, 0, 1, 1, 1, 1})
	fmt.Println("Index of first 1 in infinite binary sorted array is:", res)

	res = binarySearch.FindMinDiffElementInSortedArray([]int{1, 2, 3, 4, 7, 8, 9}, 5)
	fmt.Println("Min diff element in sorted array is found", res)

	resArr := binarySearch.FindPeakElement([]int{5, 10, 20, 90, 80})
	fmt.Printf("Peak element in array is found at index::%+v\n", resArr)

	res = binarySearch.FindElementInBitonicArray([]int{1, 3, 8, 12, 4, 2}, 4)
	fmt.Println("Element 4 found in bitonic array at index:", res)

	row, col := binarySearch.FindElementInRowColumnSortedMatrix([][]int{
		{10, 20, 30, 40},
		{15, 25, 35, 45},
		{27, 29, 37, 48},
		{32, 33, 39, 50},
	}, 29)
	fmt.Println("Element 29 found in row-column sorted matrix", row, col)

	res = binarySearch.AllocateMinNoOfPages([]int{10, 20, 30, 40}, 2)
	fmt.Println("Minimum number of pages allocated is:", res)
	//binary search : end

	// dynamic programming on trees : start
	root := &baseDS.TreeNode{Value: 1,
		Left: &baseDS.TreeNode{Value: 2,
			Left:  &baseDS.TreeNode{Value: 4},
			Right: &baseDS.TreeNode{Value: 5}},
		Right: &baseDS.TreeNode{Value: 3}}
	// Left:  &baseDS.TreeNode{Value: 6},
	// Right: &baseDS.TreeNode{Value: 7}}

	resDiameter := 0
	res = dynamicProg.DiameterOfBinaryTree(root, &resDiameter)
	fmt.Println("Diameter of binary tree is:", resDiameter, res)

	root1 := &baseDS.TreeNode{Value: 1,
		Left: &baseDS.TreeNode{Value: 2,
			Left:  &baseDS.TreeNode{Value: 4},
			Right: &baseDS.TreeNode{Value: -5}},
		Right: &baseDS.TreeNode{Value: 3,
			Left:  &baseDS.TreeNode{Value: 6},
			Right: &baseDS.TreeNode{Value: 7}}}
	res = dynamicProg.MaximumPathSum(root)
	fmt.Println("Maximum path sum in binary tree is:", res, root1.Value)
	// dynamic programming on trees : end

	// sliding window : start
	res = slidingwindow.MaxSumSubArrayOfSizeK1([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	fmt.Println("Maximum sum of subarray of size 3 is:", res)

	arr = slidingwindow.FindFirstNegativeNumberInAllWindowOfSizeK1([]int{12, -1, -7, 8, -15, 30, 16, 28}, 3)
	fmt.Printf("First negative number in all windows of size 3 is:%+v", arr)

	res = slidingwindow.CountOccurrenceOfAnagram1("forxxorfxxrof", "for")
	fmt.Println("Count of occurrences of anagram is:", res)

	resArr1 := slidingwindow.MaxInSubArray1([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	fmt.Printf("Maximum in subarray of size 3 is:%+v", resArr1)

	resArr1 = slidingwindow.MaxInSubArray1([]int{3, 1, -3, -1, 3, 5, 8}, 3)
	fmt.Printf("Maximum in subarray of size 3 is:%+v", resArr1)

	resArr1 = slidingwindow.LargestWindowSubArrayOfSumK([]int{4, 1, 1, 1, 2, 3, 5}, 5)
	fmt.Printf("Largest window subarray of sum 5 is:%+v", resArr1)

	resStr := slidingwindow.LongestSubstringWithKUniqueCharacters("aabacbebebe", 3)
	fmt.Println("Longest substring with 3 unique characters is:", resStr)

	resStr = slidingwindow.LongestSubstringWithoutRepeatingCharacters("pwwkew")
	fmt.Println("Longest substring without repeating characters is:", resStr)

	resStr = slidingwindow.MinimumWindowSubstring("ADOBECODEBANC", "ABC")
	fmt.Println("Minimum window substring is:", resStr)

	// sliding window : end

	// graph : start
	edges := [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}}
	graphAdjList := graph.CreateGraphAdjacencyList(edges)
	fmt.Println("Graph Adjacency List:", graphAdjList)

	bfsResult := graph.BreadthFirstTraversal(graphAdjList, 1)
	fmt.Println("BFS Result starting from node 1:", bfsResult)
	dfsResult := graph.DepthFirstTraversal(graphAdjList, 1)
	fmt.Println("DFS Result starting from node 1:", dfsResult)
	dfsResultRecursive := []int{}
	visited := make(map[int]bool)
	graph.DepthFirstTraversalRecursive(graphAdjList, 1, visited, &dfsResultRecursive)
	fmt.Println("DFS Recursive Result starting from node 1:", dfsResultRecursive)

	edges1 := [][]int{{1, 2}, {1, 3}, {4, 7}, {5, 7}, {6, 7}}
	graphAdjList1 := graph.CreateGraphAdjacencyList(edges1)
	fmt.Println("Graph Adjacency List:", graphAdjList1)

	pathExists := graph.FindIfPathExists(graphAdjList1, 1, 4)
	fmt.Println("Path exists from 1 to 4:", pathExists)
	pathExists = graph.FindIfPathExists(graphAdjList1, 4, 6)
	fmt.Println("Path exists from 4 to 6:", pathExists)

	knightPos := [2]int{4, 5}
	targetPos := [2]int{1, 1}
	//n := 6 // Assuming the chessboard is 6x6, adjust indices accordingly

	// src_x, src_y := n-knightPos[0], knightPos[1]-1
	// dest_x, dest_y := n-targetPos[0], targetPos[1]-1

	steps := graph.MinStepByKnightToReachTarget(knightPos, targetPos)
	fmt.Printf("Minimum steps for knight to reach target from %v to %v is: %d\n", knightPos, targetPos, steps)

	reOrderedEdges := graph.ReorderRoutes1([][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 0)
	fmt.Println("Reordered edges to make all paths lead to destination:", reOrderedEdges)

	graphAdjList2 := graph.CreateGraphAdjacencyList([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}})
	fmt.Printf("Graph Adjacency List for cycle detection: %v\n", graphAdjList2)
	isCycleDetected := graph.IsCycleInUndirectedGraph(graphAdjList2)
	fmt.Println("Cycle exists in the graph:", isCycleDetected)

	cycleExists := graph.IsCycleInUndirectedGraphByBFS(graphAdjList2)
	fmt.Println("Cycle exists in the graph using BFS:", cycleExists)

	graphAdjList3 := graph.CreateDirectedGraphAdjacencyList([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}})
	//([][]int{{1, 2}, {2, 3}, {3, 4}, {2, 4}}) -> cycle doesnt exist
	fmt.Printf("Graph Adjacency List for cycle detection: %v\n", graphAdjList3)
	isCycleDetected = graph.IsCycleInDirectedGraphByDFS(graphAdjList3)
	fmt.Println("Cycle exists in the directed graph:", isCycleDetected)

	graphAdjList4 := graph.CreateDirectedGraphAdjacencyList([][]int{{1, 4}, {4, 5}, {5, 2}, {2, 1}, {3, 2}})
	//([][]int{{1, 2}, {2, 3}, {3, 4}, {2, 4}})-> safe nodes exist
	//([][]int{{1, 4}, {4,5}, {5,2}, {2, 1},{3,2}}) -> special case
	fmt.Printf("Graph Adjacency List for cycle detection: %v\n", graphAdjList4)
	safeNodes := graph.FindSafeNodes(graphAdjList4)
	fmt.Println("Safe nodes in the directed graph:", safeNodes)

	longestCycleLength := graph.FindLongestcycle(graphAdjList4)
	fmt.Println("Longest cycle in the directed graph:", longestCycleLength)

	graphAdjList5 := graph.CreateDirectedGraphAdjacencyList([][]int{{1, 3}, {2, 3}, {3, 4}, {3, 5}, {6, 5}, {4, 7}, {5, 7}})
	sortedOrder := graph.TopologicalSortUsingBFS(graphAdjList5)
	fmt.Println("Topological Sort using BFS:", sortedOrder)

	sortedOrder = graph.TopologicalSortUsingDFS(graphAdjList5)
	fmt.Println("Topological Sort using DFS:", sortedOrder)

	sortedOrder = graph.FindNoOfCourses(4, [][]int{{1, 0}, {2, 1}, {3, 2}})
	fmt.Println("Course order to finish all courses:", sortedOrder)

	maxColorPath := graph.FindLargestColorValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}})
	fmt.Println("Maximum color path length in the graph:", maxColorPath)

	filleDImage := graph.FloodFill([][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, 1, 1, 2)
	fmt.Println("Flood filled image:", filleDImage)

	filleDImage = graph.FloodFillBFS([][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, 1, 1, 2)
	fmt.Println("Flood filled image using BFS:", filleDImage)

	noOfIslands := graph.NoOfIslands([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 1},
		{1, 0, 0, 1, 1},
	})
	fmt.Println("Number of islands in the grid:", noOfIslands)

	minTime := graph.RottingOranges([][]int{
		{2, 1, 0, 2, 1},
		{1, 0, 1, 2, 1},
		{1, 0, 0, 2, 1},
	})

	fmt.Println("Minimum time required to rot all oranges is:", minTime)

	villageDistance := graph.GeeksVillage(5, 5, [][]byte{
		{'H', 'H', 'H', 'H', 'H'},
		{'H', 'N', 'H', 'W', 'H'},
		{'H', 'H', 'H', 'H', 'H'},
		{'H', 'W', 'H', 'N', 'H'},
		{'H', 'H', 'H', 'H', 'H'},
	})
	fmt.Println("Distance of each house from nearest well in village is:", villageDistance)

	conn := graph.MakeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}})
	fmt.Println("Minimum number of operations to connect all computers is:", conn)

	moves := graph.SnakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	})

	fmt.Println("Minimum number of moves to reach end in Snakes and Ladders is:", moves)

	// Finding shortest path using Dijkstra's algorithm
	graphAdjListWithWeights := graph.CreateGraphAdjacencyListWithWeights([][]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}, {3, 4, 1}, {2, 4, 5}})
	fmt.Println("Graph Adjacency List with weights:", graphAdjListWithWeights)
	shortestPaths, distance := graph.DijkstraAlgoWithPath(graphAdjListWithWeights, 1, 4)
	fmt.Println("Shortest paths from node 1 using Dijkstra's algorithm:", shortestPaths, distance)
	// graph : end

	//map[1:0 2:1 3:3 4:4]
}
