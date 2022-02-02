package main

import (
	"fmt"
)

func main() {
	var option, key int

	root := newTree()
	for true {
		fmt.Println("What would you like to do?")
		fmt.Println("--------------------------")
		fmt.Println("1. Insert")
		fmt.Println("2. Search")
		fmt.Println("3. Print Current Tree")
		fmt.Println("4. Quit")
		_, _ = fmt.Scanf("%d", &option)
		fmt.Println("--------------------------")
		switch option {
		case 1:
			option = 0
			fmt.Println("Enter the number you would like to insert:")
			_, err := fmt.Scanf("%d", &key)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("Inserting key with value: %d\n", key)
			root.Insert(key)
			key = 0
			fmt.Println("--------------------------")

		case 2:
			option = 0
			fmt.Println("Enter the number you would like to search:")
			_, err := fmt.Scanf("%d", &key)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("Searching for key with value: %d\n", key)

			switch {
			case root.Search(key):
				fmt.Printf("%d found!", key)
			default:
				fmt.Println("Not found")
			}

			key = 0
			fmt.Println("--------------------------")

		case 3:
			root.Print()
		case 4:
			fallthrough
		default:
			option = 0
			fmt.Println("Bye!")
			return
		}

	}
}

func newTree() *Tree {
	return &Tree{Node: &Node{Key: 5}}
}

type Tree struct {
	*Node
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int) {
	// if node is nil, create a new Node with value key
	// if node's Key is less than key parameter, run on the right
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
	// if node's Key is greater than key parameter, run on the left
	if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	// recursively loop through
	// is n.Key < key?
	//	n.Right.Search(key)
	// else
	//	n.Left.Search(key)
	// when am I done?
	// 	the node you're on is nil,
	//	we are at the bottom,
	//	and we couldn't find it

	if n == nil {
		return false
	}

	if n.Key < key {
		return n.Right.Search(key)
	}

	if n.Key > key {
		return n.Left.Search(key)
	}

	return true
}

func (n *Node) Print() {
	fmt.Println(n)
	if n == nil || (n.Left == nil && n.Right == nil) {
		return
	}

	n.Left.Print()
	n.Right.Print()
}
