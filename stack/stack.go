package stack

type Stack[T any] []T

func (s *Stack[T]) push(v T) {
	*s = append([]T{v}, (*s)...)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(*s) == 0 {
		var v T
		return v, false
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v, true
}

// func main() {
// 	s := new(Stack[int])
// 	s.push(1)
// 	s.push(2)
// 	s.push(3)
// 	fmt.Println(s)
// 	fmt.Println(s.pop())
// 	fmt.Println(s.pop())
// 	fmt.Println(s)
// }
