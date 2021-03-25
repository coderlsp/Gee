package gee

import (
	"strings"
)

type node struct {
	pattern  string  // Route to be matched. Example: /p/:lang
	part     string  // Part of the route. Example: :lang
	children []*node // Children nodes. Example: [doc,xls]
	isWild   bool    // Use wildcard to match. True if part has prefix ':' or '*'
}

// Find the first successfully matched node function used by insert.
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// match all nodes function used by search.
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 开发服务时,注册路由规则,映射handler
// 访问时,匹配路由规则,查找对应的handler
// 因此,Trie Tree 需要支持节点的插入与查询

// addRoute adds a node.
// Not concurrency-safe!
func (n *node) addRoute(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.addRoute(pattern, parts, height+1)
}

// Search node which contains the parts
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
