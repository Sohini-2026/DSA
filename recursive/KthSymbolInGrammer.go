package recursive

//     0
//   0    1
// 0   1   1   0
//0 1 1 0  1 0 0 1

func GetKthSymbolInGrammer(n, k int) int {
	if n == 1 {
		return 0
	}

	mid := (1 << (n - 1)) / 2
	if k <= mid {
		return GetKthSymbolInGrammer(n-1, k)
	} else {
		// Flip the symbol
		return 1 - GetKthSymbolInGrammer(n-1, k-mid)
	}
}
