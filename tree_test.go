package vterm

import (
	"testing"
)

func TestTreePrint(t *testing.T) {
	tree := NewNode("node0")

	node1 := NewNode("node1")
	node2 := NewNode("node2")
	node3 := NewNode("node3")
	node4 := NewNode("node4")

	node1.AddLeaf("leaf1")
	node1.AddLeaf("leaf1")
	node1.AddLeaf("leaf1")

	node1.AddNode(node4)
	node3.AddLeaf("leaf3")

	node4.AddLeaf("leaf4")
	node4.AddLeaf("leaf4")
	node4.AddLeaf("leaf4")

	node5 := NewNode("node5")
	node5.AddNode(NewNode("node6"))
	node4.AddNode(node5)

	tree.AddNode(node1)
	tree.AddNode(node2)
	node1.AddNode(node3)

	node7 := NewNode("node7")
	node2.AddNode(node7)
	node7.AddLeaf("leaf7")
	node7.AddLeaf("leaf7")
	node7.AddLeaf("leaf7")

	tree.Print()
}
