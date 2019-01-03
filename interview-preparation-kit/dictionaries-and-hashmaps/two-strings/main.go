package main

import (
	"fmt"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	skip := map[rune]struct{}{}
	for _, s := range s1 {
		if _, ok := skip[s]; ok {
			continue
		}
		for _, t := range s2 {
			if s == t {
				return "YES"
			}
		}
		skip[s] = struct{}{}
	}
	return "NO"
}

func main() {
	s1 := "hello"
	s2 := "world"

	// s1 := "h1"
	// s2 := "world"

	fmt.Println(twoStrings(s1, s2))
}
