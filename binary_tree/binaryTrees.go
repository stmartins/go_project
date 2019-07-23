package main

import "github.com/user/tree"
import "fmt"

func WalkRecurs(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkRecurs(t.Left, ch)
	ch <- t.Value
	WalkRecurs(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecurs(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			fmt.Println(v1, v2)
		} else if !ok1 && !ok2 {
			return true
        }
		if v1 != v2 {
			return false
		}

	}
	fmt.Println("same function")
	return true
}

func main() {
	t1 := tree.New(2)
	t2 := tree.New(3)
	if Same(t1, t1) == true {
		fmt.Println("Same is true")
	}
	if Same(t1, t2) == false {
		fmt.Println("Same is false")
	}
}
