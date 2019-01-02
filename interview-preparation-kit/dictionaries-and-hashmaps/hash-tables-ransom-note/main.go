package main

import (
	"fmt"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	for _, n := range note {
		r := false
		for j, m := range magazine {
			if n == m {
				r = true
				magazine = append(magazine[:j], magazine[j+1:]...)
				break
			}
		}
		if r == false {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func main() {
	// magazine := []string{"two", "times", "three", "is", "not", "four"}
	// note := []string{"two", "times", "two", "is", "four"}

	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	checkMagazine(magazine, note)
}
