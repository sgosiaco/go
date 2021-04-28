package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	for _, word := range strings.Fields(s) {
		_, ok := m[word]
		if ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
}
