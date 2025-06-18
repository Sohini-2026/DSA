package recursive

func JosepheusProblem(n, k int) int {
	if n == 1 {
		return 0 // Base case: only one person left
	}
	// Recursive case: (Josephus(n-1, k) + k) % n
	return (JosepheusProblem(n-1, k) + k) % n
}

/*
Start at 1, count 3: 3 is eliminated
Remaining: [1, 2, 4, 5, 6, 7]
Next count from 4: 6 is eliminated
Remaining: [1, 2, 4, 5, 7]
Next count from 7: 2 is eliminated
Remaining: [1, 4, 5, 7]
Next count from 4: 7 is eliminated
Remaining: [1, 4, 5]
Next count from 1: 5 is eliminated
Remaining: [1, 4]
Next count from 1: 1 is eliminated
Remaining: [4] (survivor)

Output at each stage (1-based positions):
n	Eliminated	Remaining
7	3	1 2 4 5 6 7
6	6	1 2 4 5 7
5	2	1 4 5 7
4	7	1 4 5
3	5	1 4
2	1	4
Survivor: 4

Summary:
The output at each stage is the person eliminated, and the survivor is position 4.
If you want to print the elimination order, you need to simulate the process with a list, not just recursion.
*/
