package ds

// Node defines a each node in the stack
type Node struct {
	data uint8
	next *Node
}

// Stack defines a stack
type Stack struct {
	Head *Node
	size int
}

// NewStack returns a new instance of Stack
func NewStack() Stack {
	return Stack{}
}

// Push pushes a data into the stack
func (s *Stack) Push(data uint8) {
	n := Node{
		data: data,
		next: s.Head,
	}
	s.size++
	s.Head = &n
}

// Pop pops the head from the stack
func (s *Stack) Pop() *uint8 {
	if s.size == 0 {
		return nil
	}
	n := s.Head
	s.size--
	s.Head = s.Head.next
	return &n.data
}

// Size returns the size of the stack
func (s Stack) Size() int {
	return s.size
}
