package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestorSet(p *Node, q *Node) *Node {
	set := make(map[int]bool)
	for p != nil { // тут аккуратно
		set[p.Val] = true
		p = p.Parent
	}
	for q != nil {
		if set[q.Val] {
			return q
		}
		q = q.Parent
	}
	return nil
}

func lowestCommonAncestorTwoPointers(p *Node, q *Node) *Node {
	a, b := p, q
	for a != b {
		if a.Parent == nil {
			a = q
		} else {
			a = a.Parent
		}
		if b.Parent == nil {
			b = p
		} else {
			b = b.Parent
		}
	}
	return a
}
