package main

import (
	"regexp"
	"strings"
)

// самое быстрое по времени, без регулярок, просто strings.Split по "/" и скип "" и "."
func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	var stack []string
	for _, d := range dirs {
		if d == "." || d == "" {
			continue
		}
		if d == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, d)
	}

	var builder strings.Builder
	builder.WriteString("/")

	for i, dir := range stack {
		builder.WriteString(dir)
		if i != len(stack)-1 {
			builder.WriteString("/")
		}
	}
	return builder.String()
}

// strings.ReplaceAll(s, old, new string) string не заменит все вхождения на единственное
// Используй регулярки
func simplifyPath2(path string) string {
	path = strings.Trim(path, "/")

	re := regexp.MustCompile(`/+`) // O(N)
	path = re.ReplaceAllString(path, "/")

	dirs := strings.Split(path, "/")
	var stack []string
	for _, d := range dirs {
		if d == "." {
			continue
		}
		if d == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, d)
	}

	var builder strings.Builder
	builder.WriteString("/")

	for i, dir := range stack {
		builder.WriteString(dir)
		if i != len(stack)-1 {
			builder.WriteString("/")
		}
	}
	return builder.String()
}
