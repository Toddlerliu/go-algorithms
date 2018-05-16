package stack

import "errors"

// LIFO
type Stack struct {
	Ele []interface{}
}

func NewStacke() *Stack {
	return &Stack{
		Ele: make([]interface{}, 0),
	}
}

func (s Stack) Empty() bool {
	if s.Ele == nil || len(s.Ele) == 0 {
		return true
	}
	return false
}

func (s Stack) Size() int {
	return len(s.Ele)
}

// 入栈
func (s *Stack) Push(item interface{}) {
	s.Ele = append(s.Ele, item)
}

// 出栈并删除
func (s *Stack) Pop() (interface{}, error) {
	if s.Size() > 0 {
		v := s.Ele[s.Size()-1]
		s.Ele = s.Ele[:s.Size()-1] // [ )
		return v, nil
	}
	return nil, errors.New("stack empty!")
}

// 出栈不删除
func (s *Stack) Peek() interface{} {
	if s.Size() > 0 {
		return s.Ele[s.Size()-1]
	}
	return nil
}

func (s *Stack) GetAll() []interface{} {
	return s.Ele
}
