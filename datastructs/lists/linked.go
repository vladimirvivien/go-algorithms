package lists

type Node struct {
	Data int
	Next *Node
}

type Linked struct {
	Head *Node
}

func NewLinked() *Linked {
	return new(Linked)
}

func (l *Linked) Insert(data int) *Linked {
	node := &Node{Data: data}
	if l.Head == nil {
		l.Head = node
		return l
	}

	// traversal
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node

	return l
}

func (l *Linked) AsSlice() []int {
	var result []int

	// traverse
	node := l.Head
	for {
		if node != nil {
			result = append(result, node.Data)
		}
		if node.Next != nil {
			node = node.Next
		}else{
			break
		}
	}

	return result
}
