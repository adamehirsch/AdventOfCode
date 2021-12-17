package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/spf13/cobra"
)

var day2112Cmd = &cobra.Command{
	Use:                   "day2112",
	Short:                 "2021 Advent of Code Day 12",
	DisableFlagsInUseLine: true,
	Run:                   day2112Func,
}

func init() {
	rootCmd.AddCommand(day2112Cmd)
}

// Node element to keep element and next node together
type Node struct {
	value string
	isBig bool
}

// Graph is the structure that contains nodes and edges
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

// NewGraph returns a new empty graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make([]*Node, 0),
		edges: make(map[Node][]*Node),
	}
}

// AddNode inserts a new node in the graph
func (g *Graph) AddNode(el string) *Node {
	n := &Node{value: el, isBig: utils.IsUpper(el)}
	g.nodes = append(g.nodes, n)
	return n
}

// AddEdge inserts a new edge in the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Graph) RemoveNode(el string) {
	n, exists, index := NodesContains(el, g.nodes)

	if exists {
		g.nodes = append(g.nodes[:index], g.nodes[index+1:]...)
	}
	delete(g.edges, *n)
	for k, v := range g.edges {
		_, exists, index := NodesContains(el, v)
		if exists {
			g.edges[k] = append(g.edges[k][:index], g.edges[k][index+1:]...)
		}
	}
}

// String returns a string reperesentation of the node
func (n Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// String returns a string representation of the graph
func (g Graph) String() string {
	sb := strings.Builder{}
	for _, v := range g.nodes {
		sb.WriteString(v.String())
		sb.WriteString(" -> [ ")
		neighbors := g.edges[*v]
		for _, u := range neighbors {
			sb.WriteString(u.String())
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	return sb.String()
}

func NodesContains(s string, n []*Node) (*Node, bool, int) {
	for i, no := range n {
		if no.value == s {
			return no, true, i
		}
	}
	return nil, false, -1
}

func getCaves(f string) *Graph {
	file, err := utils.Opener(f, true)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	g := NewGraph()

	for _, line := range strings.Split(file, "\n") {
		nodenames := strings.Split(line, "-")
		nodes := []*Node{}
		for _, n := range nodenames {
			node, exists, _ := NodesContains(n, g.nodes)
			if !exists {
				node = g.AddNode(n)
			}
			nodes = append(nodes, node)
		}
		g.AddEdge(nodes[0], nodes[1])

	}
	return g

}

func day2112Func(cmd *cobra.Command, args []string) {
	caves := getCaves("data/2112-sample.txt")
	fmt.Print(caves)
}
