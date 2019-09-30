/*
Given a number z ∈ N and a function f(x, y) ∈ N, where x ∈ N and y ∈ N.

Function f(x, y) is black box and unknown but we can pass any x and y to function and get f(x, y) value. f(x, y) is strictly increasing:
f(x, y) < f(x + 1, y)
f(x, y) < f(x, y + 1)

Find all pairs of x and y, where f(x, y) = z.

Example 1:

Input:
f(x, y) = x + y
z = 5

Output:
[[1, 4], [2, 3], [3, 2], [4, 1]]
Example 2:

Input:
f(x, y) = x^2 + y
z = 50

Output:
[[1, 49], [2, 46], [3, 41], [4, 34], [5, 25], [6, 14], [7, 1]]
*/
package quiz

import "fmt"

func FindFuncArgumentNaive(f func(int, int) int, z int) [][]int {
	res := [][]int{}
	for x := 1; x < z; x++ {
		for y := 1; y < z-x+1; y++ {
			val := f(x, y)
			if val > z {
				break
			} else if val == z {
				res = append(res, []int{x, y})
				break
			}
		}
	}
	return res
}

func binSearch(f func(int, int) int, x int, y_min int, y_max int, z int) int {
	if y_min > y_max {
		return -1
	}
	y_mid := (y_min + y_max) / 2
	val := f(x, y_mid)
	if val == z {
		return y_mid
	} else if val > z {
		return binSearch(f, x, y_min, y_mid-1, z)
	} else {
		return binSearch(f, x, y_mid+1, y_max, z)
	}
}

func FindFuncArgumentBinary(f func(int, int) int, z int) [][]int {
	res := [][]int{}
	for x := 1; x < z; x++ {
		val := binSearch(f, x, 1, z-x+1, z)
		if val < 0 {
			break
		} else {
			res = append(res, []int{x, val})
		}
	}
	return res
}

func f1(x int, y int) int {
	return x + y
}

func f2(x int, y int) int {
	return x*x + y
}

func runFindFuncArgument() error {
	fmt.Println("Input:")
	fmt.Println("f(x, y) = x + y")
	fmt.Println("z = 5")
	fmt.Println("\nOutput:")
	fmt.Println(FindFuncArgumentBinary(f1, 5))
	fmt.Println("\nInput:")
	fmt.Println("f(x, y) = x^2 + y")
	fmt.Println("z = 50")
	fmt.Println("\nOutput:")
	fmt.Println(FindFuncArgumentBinary(f2, 50))
	return nil
}
