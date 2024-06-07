package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 0)
	a = append(a, 18)
	a = append(a, 0)
	a = append(a, 9)
	a = append(a, 18)

	fmt.Println(findOriginalArray(a))
}

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return []int{}
	}

	sort.Ints(changed)

	result := []int{}
	mp := make(map[int]int)

	for _, v := range changed {
		mp[v]++
	}

	for _, v := range changed {

		if mp[v] > 0 {
			mp[v]--
			if v, ok := mp[v*2]; ok {
				if v > 0 {
					result = append(result, v)
					mp[v*2]--
				} else {
					return []int{}
				}
			}
		}
	}
	return result
}
