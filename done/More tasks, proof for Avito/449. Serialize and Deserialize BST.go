package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor3() Codec {
	return Codec{}
}

func stringToIntSlice(s string) ([]int, error) {
	strs := strings.Split(s, " ")
	ints := make([]int, len(strs))
	for i, str := range strs {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}
	return ints, nil
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var traverse func(*TreeNode, *strings.Builder)
	traverse = func(n *TreeNode, builder *strings.Builder) {
		if n == nil {
			builder.WriteString("nil")
			builder.WriteString(" ")
			return
		}
		builder.WriteString(strconv.Itoa(n.Val))
		builder.WriteString(" ")

		traverse(n.Left, builder)
		traverse(n.Right, builder)
	}
	var builder strings.Builder
	traverse(root, &builder)
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// тут я совершила ошибку, думала что reslicing внутри функции модифицирует слайс в outer-функции.
	// https://chatgpt.com/c/671774b8-70b8-8000-89f2-05db8cd8be38
	var traverse func([]string) (*TreeNode, []string)
	traverse = func(vals []string) (*TreeNode, []string) {
		val := vals[0]
		vals = vals[1:]
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, vals
		}
		node := &TreeNode{num, nil, nil}
		node.Left, vals = traverse(vals)
		node.Right, vals = traverse(vals)
		return node, vals
	}

	vals := strings.Split(data, " ")
	root, _ := traverse(vals)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
