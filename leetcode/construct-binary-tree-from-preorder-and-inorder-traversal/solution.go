package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) ToList() []int {
	type node struct {
		n *TreeNode
		i int
		d int
	}
	stack := []node{{n, 0, 0}}
	var list []int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// if wantlen := 2*top.d + 2; len(list) < wantlen {
		if wantlen := top.i + 1; len(list) < wantlen {
			oldlen := len(list)
			list = append(list, make([]int, wantlen-oldlen)...)
			for i := oldlen; i < len(list); i++ {
				list[i] = -666
			}
		}
		list[top.i] = top.n.Val
		if top.n.Left != nil {
			stack = append(stack, node{top.n.Left, 2*top.i + 1, top.d + 1})
		}
		if top.n.Right != nil {
			stack = append(stack, node{top.n.Right, 2*top.i + 2, top.d + 1})
		}
	}
	return list
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n, _ := buildTreeImpl(preorder, inorder)
	return n
}

func buildTreeImpl(preorder []int, inorder []int) (*TreeNode, int) {
	// fmt.Println(preorder, inorder)
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil, 0
	}
	root := preorder[0]
	rootPosIO := -1
	for i, v := range inorder {
		if v == root {
			rootPosIO = i
			break
		}
	}
	// fmt.Printf("left of %v\n", root)
	leftTree, numLeft := buildTreeImpl(preorder[1:], inorder[:rootPosIO])
	// fmt.Printf("right of %v\n", root)
	rightTree, numRight := buildTreeImpl(preorder[1+numLeft:], inorder[rootPosIO+1:])
	n := &TreeNode{
		Val:   root,
		Left:  leftTree,
		Right: rightTree,
	}
	return n, 1 + numLeft + numRight
}
