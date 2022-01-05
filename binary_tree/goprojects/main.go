package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//    				tr
// 			trl				trr
// 		trl		trr		trl		trr
//	r		r		r		r		r
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

// func treeToSlice(t * Tree) []int {

// }

func create(n ...int) *Tree {
	var t *Tree
	// rand.Seed(time.Now().Unix())
	// for i := 0; i < 2*n; i++ {
	// 	temp := rand.Intn(n * 2)
	// 	t = insert(t, temp)
	// }
	for _, v := range n {
		t = insert(t, v)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create([]int{10, 3, 4, 2, 9, 5}...)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
