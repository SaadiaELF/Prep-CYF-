package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
