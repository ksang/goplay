/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/
package quiz

import "fmt"

func rotateImg(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
	for y := 0; y < n-1; y++ {
		for x := y; x < n-y-1; x++ {
			ix := x
			iy := y
			new_ix := n - iy - 1
			new_iy := ix
			tmp := matrix[new_iy][new_ix]
			matrix[new_iy][new_ix] = matrix[iy][ix]
			for i := 0; i < 3; i++ {
				ix = new_ix
				iy = new_iy
				new_ix = n - iy - 1
				new_iy = ix
				tmptmp := matrix[new_iy][new_ix]
				matrix[new_iy][new_ix] = tmp
				tmp = tmptmp
			}
		}
	}

}

func printImg(matrix [][]int) {
	fmt.Println("[")
	for _, v := range matrix {
		fmt.Printf("    %v\n", v)
	}
	fmt.Println("]")
}

func runRotateImg() error {
	matrix1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println("Input:")
	printImg(matrix1)
	rotateImg(matrix1)
	fmt.Println("Output:")
	printImg(matrix1)
	matrix2 := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	fmt.Println("Input:")
	printImg(matrix2)
	rotateImg(matrix2)
	fmt.Println("Output:")
	printImg(matrix2)
	return nil
}
