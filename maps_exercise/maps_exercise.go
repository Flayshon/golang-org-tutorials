package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	
	for _, value := range words {
		m[value]++
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
