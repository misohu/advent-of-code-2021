package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	left, right int
	parent      *Node
	leftNode    *Node
	rightNode   *Node
}

func main() {
	var err error
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %v: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()
	numbers := readFile(file)
	res := numbers[0]
	var node *Node
	for _, s := range numbers[1:] {
		newSum := "[" + res + "," + s + "]"
		fmt.Println(newSum)
		node, _, err = parseNode(newSum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %v: %v\n", newSum, err)
			os.Exit(1)
		}
		processSum(node)
		res = printTree(node)
	}
	fmt.Println(res)
	fmt.Println(contMagnitude(node))
	// tmp, _, err := parseNode("[9,1]")
	// fmt.Println(contMagnitude(tmp))
}

func contMagnitude(n *Node) int {
	res := 0
	if n.leftNode != nil {
		res += 3 * contMagnitude(n.leftNode)
	} else {
		res += 3 * n.left
	}
	if n.rightNode != nil {
		res += 2 * contMagnitude(n.rightNode)
	} else {
		res += 2 * n.right
	}
	return res
}

func readFile(f *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func processSum(n *Node) {
	fmt.Printf("RECAIVED: ")
	fmt.Println(printTree(n))
	for {
		t := findAtDepth(n, 0, 5)
		if t != nil {
			l := findLeftParent(t, t.left)
			r := findRightParent(t, t.right)
			fmt.Printf("FOUND: %v, %v\n", l, r)
			if t.parent.leftNode == t {
				t.parent.left = 0
				t.parent.leftNode = nil
			} else {
				t.parent.right = 0
				t.parent.rightNode = nil
			}
			fmt.Printf("EXPLODED: ")
			fmt.Println(printTree(n))
			continue
		}
		if split(n) {
			fmt.Printf("SPLIT: ")
			fmt.Println(printTree(n))
			continue
		}
		break
	}
}

func findRightParent(n *Node, val int) int {
	if n.parent == nil {
		return -1
	}
	if n.parent.right != -1 {
		n.parent.right += val
		return n.parent.right
	}
	if n.parent.rightNode == n {
		return findRightParent(n.parent, val)
	}
	return findLeft(n.parent.rightNode, val)
}

func findLeftParent(n *Node, val int) int {
	fmt.Println(n)
	if n.parent == nil {
		return -1
	}
	if n.parent.left != -1 {
		n.parent.left += val
		return n.parent.left
	}
	if n.parent.leftNode == n {
		return findLeftParent(n.parent, val)
	}
	return findRight(n.parent.leftNode, val)
}

func findLeft(n *Node, val int) int {
	if n.leftNode == nil {
		n.left += val
		return n.left
	}
	return findLeft(n.leftNode, val)
}

func findRight(n *Node, val int) int {
	if n.rightNode == nil {
		n.right += val
		return n.right
	}
	return findRight(n.rightNode, val)
}

func findAtDepth(n *Node, curr, target int) *Node {
	curr += 1
	if curr == target {
		return n
	}
	if n.leftNode != nil {
		if node := findAtDepth(n.leftNode, curr, target); node != nil {
			return node
		}
	}
	if n.rightNode != nil {
		if node := findAtDepth(n.rightNode, curr, target); node != nil {
			return node
		}
	}
	return nil
}

func parseNode(s string) (*Node, string, error) {
	node := &Node{}
	var err error
	if s[0] != '[' {
		return nil, "", fmt.Errorf("parseNode: expected '[' got %v", s)
	}
	s = s[1:]
	if s[0] == '[' {
		node.leftNode, s, err = parseNode(s)
		if err != nil {
			return nil, "", fmt.Errorf("parseNode: %v %v", s, err)
		}
		node.left = -1
		node.leftNode.parent = node
	} else {
		node.left, err = strconv.Atoi(string(s[0]))
		if err != nil {
			return nil, "", fmt.Errorf("parseNode: %v %v", s, err)
		}
		s = s[1:]
	}

	if s[0] == ',' {
		s = s[1:]
	} else {
		return nil, "", fmt.Errorf("parseNode: Expected comma %v", s)
	}
	if s[0] == '[' {
		node.rightNode, s, err = parseNode(s)
		if err != nil {
			return nil, "", fmt.Errorf("parseNode: %v %v", s, err)
		}
		node.right = -1
		node.rightNode.parent = node
	} else {
		node.right, err = strconv.Atoi(string(s[0]))
		if err != nil {
			return nil, "", fmt.Errorf("parseNode: %v %v", s, err)
		}
		s = s[1:]
	}
	if s[0] == ']' {
		s = s[1:]
	} else {
		return nil, "", fmt.Errorf("parseNode: Expected closing bracket %v", s)
	}
	return node, s, nil
}

func printTree(n *Node) string {
	s := "["
	if n.leftNode != nil {
		s += printTree(n.leftNode)
	} else {
		s += fmt.Sprintf("%v", n.left)
	}
	s += ","
	if n.rightNode != nil {
		s += printTree(n.rightNode)
	} else {
		s += fmt.Sprintf("%v", n.right)
	}
	s += "]"
	return s
}

func split(n *Node) bool {
	if n.left != -1 {
		if n.left > 9 {
			n.leftNode = &Node{
				right:  int(math.Ceil(float64(n.left) / 2)),
				left:   int(math.Floor(float64(n.left) / 2)),
				parent: n}
			n.left = -1
			return true
		}
	}
	if n.leftNode != nil {
		if split(n.leftNode) {
			return true
		}
	}
	if n.right != -1 {
		if n.right > 9 {
			n.rightNode = &Node{
				right:  int(math.Ceil(float64(n.right) / 2)),
				left:   int(math.Floor(float64(n.right) / 2)),
				parent: n}
			n.right = -1
			return true
		}
	}
	if n.rightNode != nil {
		if split(n.rightNode) {
			return true
		}
	}
	return false
}
