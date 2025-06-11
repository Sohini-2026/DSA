package recursive

func TowerOfHanoi(n int, source, destination, auxiliary string) {
	if n == 1 {
		println("Move disk 1 from", source, "to", destination)
		return
	}

	TowerOfHanoi(n-1, source, auxiliary, destination)
	println("Move disk", n, "from", source, "to", destination)
	TowerOfHanoi(n-1, auxiliary, destination, source)
}

// Example usage:
// func main() {
// 	n := 3
// 	TowerOfHanoi(n, "A", "C", "B")
// }
// Output:
// Move disk 1 from A to C
// Move disk 2 from A to B
// Move disk 1 from C to B
// Move disk 3 from A to C
