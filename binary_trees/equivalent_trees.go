package main

import "golang.org/x/tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
        Walk(t.Left, ch)
        ch <- t.Value
        Walk(t.Right, ch)
    }
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(7)
	t2 := tree.New(7)

	fmt.Println(Same(t1, t2))
}
