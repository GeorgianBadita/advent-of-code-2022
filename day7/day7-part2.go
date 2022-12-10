package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Name     string
	Children []*TreeNode
	IsDir    bool
	Size     int
	Parent   *TreeNode
}

func (d *TreeNode) AppendChild(nd *TreeNode) *TreeNode {
	for _, cnd := range d.Children {
		if nd.Name == cnd.Name && nd.IsDir == cnd.IsDir {
			return cnd
		}
	}
	(*d).Children = append((*d).Children, nd)
	return nd
}

func readFromFile(file string) (*TreeNode, error) {
	fileH, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fileH)

	root := &TreeNode{Name: "/", Children: []*TreeNode{}, IsDir: true, Size: 0, Parent: nil}
	currNode := root

	// Skip first line as it is a cd /
	if scanner.Scan() {
		scanner.Text()
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			// We have a command
			line = line[2:]
			if line[:2] == "cd" {
				if line[3] == '/' {
					currNode = root
				} else if line[3:] == ".." {
					nextNd := currNode.Parent
					currNode = nextNd
				} else {
					newDir := &TreeNode{IsDir: true, Name: line[3:], Children: []*TreeNode{}, Parent: currNode}
					currNode = currNode.AppendChild(newDir)
				}
			} else {
				if line[:2] == "ls" {
					continue
				}
			}
		} else {
			// We have output
			lsOutput := strings.Split(line, " ")
			if lsOutput[0] == "dir" {
				newDir := &TreeNode{Name: lsOutput[1], IsDir: true, Children: []*TreeNode{}, Parent: currNode}
				currNode.AppendChild(newDir)
			} else {
				conv, _ := strconv.Atoi(lsOutput[0])
				newFile := &TreeNode{Name: lsOutput[1], IsDir: false, Size: conv, Parent: currNode}
				currNode.AppendChild(newFile)
			}
		}
	}
	return root, nil
}

func ComputeSizes(node *TreeNode) int {
	if !node.IsDir {
		return node.Size
	}

	sz := 0
	for _, nd := range node.Children {
		sz += ComputeSizes(nd)
	}

	node.Size = sz
	return node.Size
}

func Solve(root *TreeNode, neededSize int) *TreeNode {
	st := []*TreeNode{root}
	bestNd := root
	bestSz := root.Size - neededSize
	for len(st) > 0 {
		nd := st[len(st)-1]
		st = st[:len(st)-1]
		if nd.Size > neededSize && nd.Size-neededSize < bestSz {
			bestSz = nd.Size - neededSize
			bestNd = nd
		}

		for _, ch := range nd.Children {
			st = append(st, ch)
		}
	}
	return bestNd
}

func main() {
	root, err := readFromFile("./in-day7.txt")
	if err != nil {
		panic("Error reading from file...")
	}

	ComputeSizes(root)
	neededSize := 30000000 - (70000000 - root.Size)
	nd := Solve(root, neededSize)
	fmt.Printf("%v\n", nd.Size)
}
