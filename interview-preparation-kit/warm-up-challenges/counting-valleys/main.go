package main

import (
	"fmt"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	cnt := int32(0)
	current := int32(0)

	for i := int32(0); i < n; i++ {
		v := string(s[i])
		if v == "U" {
			current++
			if current == 0 {
				cnt++
			}
		} else {
			current--
		}
	}
	return cnt
}

func main() {
	n := int32(8)
	s := "UDDDUDUU"
	fmt.Println(countingValleys(n, s))
}
