/*
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/
package quiz

import "fmt"

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func editDistance(w1 string, w2 string) int {
	m := len(w1)
	n := len(w2)
	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func runEditDistance() error {
	w1 := "horse"
	w2 := "ros"
	fmt.Printf("Input: word1 = \"%s\", word2 = \"%s\"\n", w1, w2)
	fmt.Printf("Output: %d\n", editDistance(w1, w2))
	w1 = "intention"
	w2 = "execution"
	fmt.Printf("Input: word1 = \"%s\", word2 = \"%s\"\n", w1, w2)
	fmt.Printf("Output: %d\n", editDistance(w1, w2))
	return nil
}
