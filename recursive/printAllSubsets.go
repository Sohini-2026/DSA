package recursive

import (
	"fmt"
	"sort"
	"strconv"
)

func PrintAllSubsets(arr []int, index int, current []int) {
	if index == len(arr) {
		// Print the current subset
		sort.Ints(current) // Sort to ensure unique subsets are printed in order
		printSubset(current)
		return
	}

	// Include the current element
	PrintAllSubsets(arr, index+1, append(current, arr[index]))

	// Exclude the current element
	PrintAllSubsets(arr, index+1, current)
}

func printSubset(subset []int) {
	if len(subset) == 0 {
		println("[]")
		return
	}

	subsetStr := "["
	for i, v := range subset {
		if i > 0 {
			subsetStr += ", "
		}
		subsetStr += strconv.Itoa(v)
	}
	subsetStr += "]"
	println(subsetStr)
}

// PrintAllSubsets prints all subsets of the given array.
// It uses recursion to explore both including and excluding each element.

func PrintAllSubsets1(input, output string) {
	if len(input) == 0 {
		println(output)
		return
	}
	output1 := output
	output2 := output + string(input[0])

	input = input[1:] // Exclude the first character

	fmt.Println("Input:", input, "Output1:", output1, "Output2:", output2)
	// Include the first character
	PrintAllSubsets1(input, output1)

	// Exclude the first character
	PrintAllSubsets1(input, output2)
}

/*
PrintAllSubsets1("abc", "")
├─ Exclude 'a': PrintAllSubsets1("bc", "")
│  ├─ Exclude 'b': PrintAllSubsets1("c", "")
│  │  ├─ Exclude 'c': PrintAllSubsets1("", "") → prints ""
│  │  └─ Include 'c': PrintAllSubsets1("", "c") → prints "c"
│  └─ Include 'b': PrintAllSubsets1("c", "b")
│     ├─ Exclude 'c': PrintAllSubsets1("", "b") → prints "b"
│     └─ Include 'c': PrintAllSubsets1("", "bc") → prints "bc"
└─ Include 'a': PrintAllSubsets1("bc", "a")
   ├─ Exclude 'b': PrintAllSubsets1("c", "a")
   │  ├─ Exclude 'c': PrintAllSubsets1("", "a") → prints "a"
   │  └─ Include 'c': PrintAllSubsets1("", "ac") → prints "ac"
   └─ Include 'b': PrintAllSubsets1("c", "ab")
      ├─ Exclude 'c': PrintAllSubsets1("", "ab") → prints "ab"
      └─ Include 'c': PrintAllSubsets1("", "abc") → prints "abc"
*/
/*
PrintAllSubsets1("abc", "")
├─ Exclude 'a'
│   PrintAllSubsets1("bc", "")
│   ├─ Exclude 'b'
│   │   PrintAllSubsets1("c", "")
│   │   ├─ Exclude 'c'
│   │   │   PrintAllSubsets1("", "") → prints ""
│   │   └─ Include 'c'
│   │       PrintAllSubsets1("", "c") → prints "c"
│   └─ Include 'b'
│       PrintAllSubsets1("c", "b")
│       ├─ Exclude 'c'
│       │   PrintAllSubsets1("", "b") → prints "b"
│       └─ Include 'c'
│           PrintAllSubsets1("", "bc") → prints "bc"
└─ Include 'a'
    PrintAllSubsets1("bc", "a")
    ├─ Exclude 'b'
    │   PrintAllSubsets1("c", "a")
    │   ├─ Exclude 'c'
    │   │   PrintAllSubsets1("", "a") → prints "a"
    │   └─ Include 'c'
    │       PrintAllSubsets1("", "ac") → prints "ac"
    └─ Include 'b'
        PrintAllSubsets1("c", "ab")
        ├─ Exclude 'c'
        │   PrintAllSubsets1("", "ab") → prints "ab"
        └─ Include 'c'
            PrintAllSubsets1("", "abc") → prints "abc"
*/
