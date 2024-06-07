package main

import (
	"fmt"
)

func main() {
	a := "dslfr"
	fmt.Println(checkIfPangram(a))
}

func checkIfPangram(sentence string) bool {
	mp := make(map[rune]struct{})
	for _, v := range sentence {
		if _, ok := mp[v]; !ok {
			mp[v] = struct{}{}
		}
	}
	return len(mp) == 26
}
