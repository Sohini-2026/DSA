package recursive

func SortArray(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}

	n := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]

	SortArray(arr)

	Insert(arr, n)
}

func Insert(arr *[]int, n int) {
	if len(*arr) == 0 || (*arr)[len(*arr)-1] <= n {
		*arr = append(*arr, n)
		return
	}

	temp := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	Insert(arr, n)
	*arr = append(*arr, temp)
}
