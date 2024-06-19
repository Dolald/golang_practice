package main

import "fmt"

func main() {
	a := "asdfgrs"

	fmt.Println(lengthOfLongestSubstring(a))
}

func lengthOfLongestSubstring(s string) int {
	var result, left int
	mp := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		mp[s[right]]++

		for mp[s[right]] > 1 {
			mp[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}
