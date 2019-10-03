package quiz

import (
	"reflect"
	"testing"
)

func TestCreateBinarySearchTree(t *testing.T) {
	tests := []struct {
		valueList []int
		expected  []int
	}{
		{
			[]int{5, 3, 2, 6, 1, 7, 99},
			[]int{1, 2, 3, 5, 6, 7, 99},
		},
		{
			[]int{7, 3, 2, 1, 3, 99},
			[]int{1, 2, 3, 3, 7, 99},
		},
	}
	for idx, c := range tests {
		root := createBinarySearchTree(c.valueList)
		eval := traverseBinaryTree(root)
		if !reflect.DeepEqual(c.expected, eval) {
			t.Errorf("case: #%d expected: %v actual: %v", idx+1, c.valueList, eval)
		}

	}
}
