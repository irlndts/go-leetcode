package leetcode

import "strings"

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	dir := NewDir()
	for _, part := range parts {
		dir.Append(part)
	}

	var result []string
	for dir.Tail != nil {
		result = append(result, dir.Pop())
	}

	output := ""
	for i := len(result) - 1; i >= 0; i-- {
		output += "/" + result[i]
	}

	if output == "" {
		return "/"
	}

	return output
}

type Dir struct {
	Tail *Node
}

type Node struct {
	Val  string
	Left *Node
}

func NewDir() *Dir {
	return &Dir{}
}

func (d *Dir) Append(val string) {
	if val == "" {
		return
	}

	if val == "." {
		return
	}

	if val == ".." {
		d.Pop()
		return
	}

	tail := d.Tail
	d.Tail = &Node{val, tail}
}

func (d *Dir) Pop() string {
	if d.Tail == nil {
		return ""
	}
	item := d.Tail.Val
	d.Tail = d.Tail.Left
	return item
}
