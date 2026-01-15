// A struct for making workig with stacks nice and easy.
type stack[T any] struct {
	data []T
}

func (s *stack[T]) push(x T) {
	s.data = append(s.data, x)
}

// Note: no safety on peek and pop
func (s *stack[T]) peek() T {
	return s.data[len(s.data)-1]
}

func (s *stack[T]) pop() T {
	ret := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return ret
}