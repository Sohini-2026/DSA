package slidingwindow

func CountOccurrenceOfAnagram(s string, p string) int {
	if len(s) < len(p) {
		return 0
	}

	countP := make(map[byte]int)
	countS := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		countP[p[i]]++
		countS[s[i]]++
	}

	result := 0
	if isAnagram(countP, countS) {
		result++
	}

	for i := len(p); i < len(s); i++ {
		countS[s[i]]++
		countS[s[i-len(p)]]--

		if isAnagram(countP, countS) {
			result++
		}
	}

	return result
}

func isAnagram(countP map[byte]int, countS map[byte]int) bool {
	if len(countP) != len(countS) {
		return false
	}

	for key, val := range countP {
		if countS[key] != val {
			return false
		}
	}

	return true
}

func CountOccurrenceOfAnagram1(s string, p string) int {
	if len(s) < len(p) {
		return 0
	}

	k := len(p)
	countP := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		countP[p[i]]++
	}

	i, j, res := 0, 0, 0
	count := len(countP) // Number of unique characters in p

	for j < len(s) {
		if _, ok := countP[s[j]]; ok {
			countP[s[j]]--
			if countP[s[j]] == 0 {
				count--
			}
		}
		if j-i+1 == k {
			if count == 0 { // If all characters in countP are zero, we have an anagram
				res++
			}
			// Slide the window
			if _, ok := countP[s[i]]; ok {
				if countP[s[i]] == 0 {
					count++ //Its becoz count is for the unique characters not frequency
				}
				countP[s[i]]++ // Restore the character that was removed from the window
			}
			i++
		}
		j++ // Expand the window
	}

	return res
}

/*
Step-by-Step Execution
1. Initialization
k = 3
countP = {'f':1, 'o':1, 'r':1}
count = 3 (unique characters in p)
i = 0, j = 0, res = 0
2. Sliding Window
First window: "for" (i=0, j=2)
j=0: s[0]='f', countP['f'] becomes 0, count=2
j=1: s[1]='o', countP['o'] becomes 0, count=1
j=2: s[2]='r', countP['r'] becomes 0, count=0
Now, window size is 3 (j-i+1 == k), and count == 0
→ Found an anagram, res = 1

Slide window:

s[i]='f', countP['f'] is 0, so count=1, countP['f'] becomes 1
i=1
Second window: "orx" (i=1, j=3)
j=3: s[3]='x', not in countP, nothing changes
Window size is 3, but count != 0
→ Not an anagram

Slide window:

s[i]='o', countP['o'] is 0, so count=2, countP['o'] becomes 1
i=2
Third window: "rxx" (i=2, j=4)
j=4: s[4]='x', not in countP, nothing changes
Window size is 3, count != 0
→ Not an anagram

Slide window:

s[i]='r', countP['r'] is 0, so count=3, countP['r'] becomes 1
i=3
Fourth window: "xxo" (i=3, j=5)
j=5: s[5]='o', countP['o'] becomes 0, count=2
Window size is 3, count != 0
→ Not an anagram

Slide window:

s[i]='x', not in countP, nothing changes
i=4
Fifth window: "xor" (i=4, j=6)
j=6: s[6]='r', countP['r'] becomes 0, count=1
Window size is 3, count != 0
→ Not an anagram

Slide window:

s[i]='x', not in countP, nothing changes
i=5
Sixth window: "orf" (i=5, j=7)
j=7: s[7]='f', countP['f'] becomes 0, count=0
Window size is 3, count == 0
→ Found an anagram, res = 2

Slide window:

s[i]='o', countP['o'] is 0, so count=1, countP['o'] becomes 1
i=6
*/
