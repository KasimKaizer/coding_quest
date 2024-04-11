package tree

import "strconv"

// I would use BFS for this problem
// construct a binary tree then walk through it
// create a queue for the nodes to visit next
// count the number of nodes visited
// record the highest nextNodes, that would be our max width
//
// The data is given as hexadecimal, so we convert it into int then add it to the binary tree

type Node struct {
	Data  int
	Right *Node
	Left  *Node
}

type Tree struct {
	Head *Node
}

func NewTree(hexData []string) (*Tree, error) {
	tree := new(Tree)
	for _, hexValue := range hexData {
		num, err := strconv.ParseInt(hexValue, 16, 64)
		if err != nil {
			return nil, err
		}
		tree.Add(int(num))
	}
	return tree, nil
}

func (t *Tree) Add(item int) {
	if t.Head == nil {
		t.Head = &Node{Data: item}
		return
	}
	next := t.Head
	for next != nil {
		if item >= next.Data {
			if next.Right == nil {
				next.Right = &Node{Data: item}
				break
			}
			next = next.Right
		} else {
			if next.Left == nil {
				next.Left = &Node{Data: item}
				break
			}
			next = next.Left
		}
	}
}

func (t *Tree) MaxWidthAndHeight() (int, int) {
	if t.Head == nil {
		return 0, 0
	}
	queue := []*Node{t.Head}
	width, height := 0, 0
	curNodes, nextNodes := 1, 0
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextNodes++
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			nextNodes++
		}
		curNodes--
		if curNodes == 0 {
			height++
			width = max(width, nextNodes)
			curNodes, nextNodes = nextNodes, 0

		}
	}
	return width, height
}
