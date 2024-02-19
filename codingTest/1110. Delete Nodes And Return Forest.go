package codingTest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := make(map[int]*TreeNode)
	parents := make(map[int]int)
	preorder(&m, &parents, root, root.Val)

	for _, val := range to_delete {
		if m[val].Left != nil {
			parents[m[val].Left.Val] = m[val].Left.Val
		}
		if m[val].Right != nil {
			parents[m[val].Right.Val] = m[val].Right.Val
		}

		if m[parents[val]].Left == m[val] {
			m[parents[val]].Left = nil
		} else if m[parents[val]].Right == m[val] {
			m[parents[val]].Right = nil
		}

		parents[val] = 0
	}

	var result []*TreeNode
	for idx, parent := range parents {
		if idx == parent {
			result = append(result, m[idx])
		}
	}

	return result
}

func preorder(m *map[int]*TreeNode, parents *map[int]int, node *TreeNode, parent int) {
	(*m)[node.Val] = node
	(*parents)[node.Val] = parent

	if node.Left != nil {
		preorder(m, parents, node.Left, node.Val)
	}
	if node.Right != nil {
		preorder(m, parents, node.Right, node.Val)
	}
}
