/*
Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    1
   / \
  2   5
 / \   \
3   4   6
The flattened tree should look like:

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

*/
package quiz

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertBinarySearchTree(root *TreeNode, node *TreeNode) {
	if root == nil {
		root = node
		return
	}
	if node.Val > root.Val {
		if root.Right == nil {
			root.Right = node
			return
		} else {
			insertBinarySearchTree(root.Right, node)
			return
		}
	} else {
		if root.Left == nil {
			root.Left = node
			return
		} else {
			insertBinarySearchTree(root.Left, node)
			return
		}
	}
}

func createBinarySearchTree(valueList []int) *TreeNode {
	if len(valueList) == 0 {
		return nil
	}
	root := TreeNode{
		Val: valueList[0],
	}
	for _, n := range valueList[1:] {
		node := TreeNode{
			Val: n,
		}
		insertBinarySearchTree(&root, &node)
	}
	return &root
}

func recBinTree(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recBinTree(root.Left, res)
	*res = append(*res, root.Val)
	recBinTree(root.Right, res)
}
func traverseBinaryTree(root *TreeNode) []int {
	res := []int{}
	recBinTree(root, &res)
	return res
}

func rightestChild(root *TreeNode) *TreeNode {
	var node *TreeNode = root
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		saved := root.Right
		root.Right = root.Left
		rightestChild(root.Left).Right = saved
		root.Left = nil
	}
	flatten(root.Right)
}

func runFlattenBinaryTreeToLinkedList() error {
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
	flatten(root)
	return nil
}
