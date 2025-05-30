package tripofgo

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(value int) {
	//make newNode
	newNode := &Node{}

	// if node is empty
	if l.Head == nil {
		l.Head = newNode
		return
	}

	//current is head node/ new node
	current := l.Head

	//iterate all node
	for current.Next != nil {
		current = current.Next
	}

	// next node is new node
	current.Next = newNode

}
