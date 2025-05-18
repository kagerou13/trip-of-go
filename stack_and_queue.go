package tripofgo

type Stack struct {
	element []int
}

func (s *Stack) Push(n int) {
	s.element = append(s.element, n)
}

func (s *Stack) Pop() int {
	//if stack is empty
	if len(s.element) == 0 {
		panic("Empty stack")
	}

	//get last element from slice
	top := s.element[len(s.element)-1]
	// delete last element and update current element
	s.element = s.element[:len(s.element)-1]

	return top
}

func (s *Stack) Peek() int {
	// if stack is empty
	if len(s.element) == 0 {
		panic("Empty stack")
	}

	// get last element without delete
	top := s.element[len(s.element)-1]

	return top
}
