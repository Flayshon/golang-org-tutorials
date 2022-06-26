package main

import (
	"errors"
	"fmt"
	"os"
)

type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(v T) {
	if l.next == nil {
		l.next = &List[T]{val: v}
	} else {
		l.next.Add(v)
	}
}

func (l *List[T]) Search(v T) (*List[T], error) {
	if l.val == v {
		return l, nil
	} else if l.next == nil {
		return nil, errors.New("the value could not be found in the list")
	} else {
		return l.next.Search(v)
	}
}

func main() {
	list := &List[int]{val: 12}
	fmt.Println(list.val)
	list.Add(23)
	list.Add(324)
	x, err := list.Search(3243)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(x.val)
}
