package main

import (
	"fmt"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Insert inserts a new value into the binary tree
func (t *TreeNode) Insert(value int) {
	if value < t.value {
		// Insert into the left subtree
		if t.left == nil {
			t.left = &TreeNode{value: value}
		} else {
			t.left.Insert(value)
		}
	} else {
		// Insert into the right subtree
		if t.right == nil {
			t.right = &TreeNode{value: value}
		} else {
			t.right.Insert(value)
		}
	}
}

// Search searches for a value in the binary tree
func (t *TreeNode) Search(value int) bool {
	if t == nil {
		return false
	}
	if value == t.value {
		return true
	} else if value < t.value {
		return t.left.Search(value)
	} else {
		return t.right.Search(value)
	}
}

// InOrderTraversal prints the values of the tree in-order (left, root, right)
func (t *TreeNode) InOrderTraversal() {
	if t == nil {
		return
	}
	t.left.InOrderTraversal()
	fmt.Printf("%d ", t.value)
	t.right.InOrderTraversal()
}

func main() {
	// Create a new root node
	root := &TreeNode{value: 10}

	// Insert values into the binary tree
	values := []int{5, 15, 3, 7, 13, 17}
	for _, v := range values {
		root.Insert(v)
	}

	// In-order traversal (sorted output)
	fmt.Println("In-order traversal:")
	root.InOrderTraversal()
	fmt.Println()

	// Search for values
	fmt.Println("Search for 7:", root.Search(7))  // Should return true
	fmt.Println("Search for 20:", root.Search(20)) // Should return false
}
