package vterm

import (
	"fmt"
	"strings"
)

type Node struct {
	name  string
	leafs []string
	nodes []*Node
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) AddLeaf(leaf string) {
	n.leafs = append(n.leafs, leaf)
}

func (n *Node) AddNode(node *Node) {
	n.nodes = append(n.nodes, node)
}

func (n *Node) Print() {
	fmt.Printf("%s\n", n.name)
	n.printNode(1, "")
}

func (n *Node) printNode(lvl int, previousIndent string) {
	// TODO: put into consts on package level
	space := "│   "

	inner := "├──"

	outer := "└──"

	nodeFmt := "%s %s\n"
	leafFmt := "%s %s\n"
	indent := ""

	lenLeafs, lenNodes := len(n.leafs), len(n.nodes)
	for i, node := range n.nodes {
		indent = previousIndent + inner
		if lenNodes-1 == i && lenLeafs == 0 {
			indent = previousIndent + outer
		}
		fmt.Printf(nodeFmt, indent, node.name)
		switch {
		case isLast(lenNodes, i) && lenLeafs == 0:
			node.printNode(lvl+1, previousIndent+"    ")
		case isLast(lenNodes, i):
			node.printNode(lvl+1, previousIndent+space)
		default:
			node.printNode(lvl+1, strings.Repeat(space, lvl))
		}
	}

	lastLeaf := lenLeafs - 1
	for i, leaf := range n.leafs {
		indent = previousIndent + inner
		if i == lastLeaf {
			indent = previousIndent + outer
		}
		fmt.Printf(leafFmt, indent, leaf)
	}
}

func isLast(lenNodes, index int) bool {
	if lenNodes-1 == index {
		return true
	}
	return false
}
