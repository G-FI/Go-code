package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string
	part     string
	childern []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern : %s, part : %s, isWild : %t\n}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {

	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.childern = append(n.childern, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {

	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	children := n.matchChildren(parts[height])
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

//matchChild 返回n的匹配到part的孩子
func (n *node) matchChild(part string) *node {
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	//只添加叶节点
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.childern {
		child.travel(list)
	}
}

//matchChildren 返回匹配到的所有parts孩子(因为占位符也在可选项中，插入的时候，不会重复)
func (n *node) matchChildren(part string) []*node {

	children := make([]*node, 0)
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return children
}
