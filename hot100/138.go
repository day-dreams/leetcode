package hot100

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var (
		memory     []*Node
		node2index = make(map[*Node]int)
		newNodes   []*Node
	)
	for current := head; current != nil; current = current.Next {
		memory = append(memory, current)
		node2index[current] = len(memory) - 1
		newNode := &Node{Val: current.Val}
		if len(newNodes) != 0 {
			newNodes[len(newNodes)-1].Next = newNode
		}
		newNodes = append(newNodes, newNode)
	}

	for i := 0; i < len(memory); i++ {
		srcNode := memory[i]
		index, ok := node2index[srcNode.Random]
		if ok {
			newNodes[i].Random = newNodes[index]
		} else {
			newNodes[i].Random = nil
		}
	}
	if len(newNodes) == 0 {
		return nil
	}
	return newNodes[0]
}
