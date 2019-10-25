/**
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

package quiz

import "fmt"

func levelOrder(root *TreeNode) [][]int {
	q := make([]*TreeNode, 0)
	res := [][]int{}
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		cr := []int{}
		nq := make([]*TreeNode, 0)
		for _, n := range q {
			cr = append(cr, n.Val)
			if n.Left != nil {
				nq = append(nq, n.Left)
			}
			if n.Right != nil {
				nq = append(nq, n.Right)
			}
		}
		q = nq
		res = append(res, cr)
	}

	return res
}

func runbinTreeLevelOrder() error {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(levelOrder(root))
	return nil
}
