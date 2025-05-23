package recursive

import "fmt"

func Print1ToN(n int) int {
	if n == 1 {
		fmt.Println(1)
		return 1
	}
	Print1ToN(n - 1)
	fmt.Println(n)
	return n
}

func PrintNTo1(n int) int {
	if n == 1 {
		fmt.Println(1)
		return 1
	}

	fmt.Println(n)
	PrintNTo1(n - 1)
	return n

}

func Fibonacci(n int) int {
	if n == 1 {
		return 1
	}

	return n * Fibonacci(n-1)
}
