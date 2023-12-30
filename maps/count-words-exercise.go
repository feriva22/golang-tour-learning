package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fmt.Println(strings.Fields(s))

	string_count := make(map[string]int)
	for _, v := range strings.Fields(s) {
		//retrieve are the word already exist, if yes append the value
		value_word, is_exist := string_count[v]
		if is_exist {
			string_count[v] = value_word + 1
		}
		if !is_exist {
			string_count[v] = 1
		}
	}
	return string_count
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
