package main

import "fmt"


func main(){
fmt.Println("What did u return ?",print1ToN(5))
printNTo1(5)
fmt.Println("Fibonacci of 5 is ",fibonacci(5))
}


func print1ToN(n int)int{
	if n==1{
		fmt.Println(1)
		return 1
	}
    print1ToN(n-1)
	fmt.Println(n)
	return n
}

func printNTo1(n int)int{
    if n==1{
		fmt.Println(1)
		return 1
	}

	fmt.Println(n)
	printNTo1(n-1)
    return n

}

func fibonacci(n int)int{
     if n==1{
          return 1
	 }

	 return n*fibonacci(n-1)
}