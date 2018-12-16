package main

import (
	"fmt"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	lastLen := n % int64(len(s))
	repeat := n / int64(len(s))

	acnt := int64(strings.Count(s, "a"))
	acnt *= repeat
	acnt += int64(strings.Count(s[:lastLen], "a"))

	return acnt
}

func main() {
	s := "aba"
	n := int64(10)

	fmt.Println(repeatedString(s, n))
}
