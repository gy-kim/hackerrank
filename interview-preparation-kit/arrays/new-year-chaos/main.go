package main

import (
	"fmt"
	"math"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {

	bribes := 0
	for i := int32(len(q) - 1); i >= 0; i-- {
		if (q[i] - (i + 1)) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		for j := int32(math.Max(0, float64(q[i]-2))); j < i; j++ {
			// fmt.Println("i", i, "j", j, "q[i]", q[i], "q[j]", q[j], "q[j] > q[i]", q[j] > q[i])
			if q[j] > q[i] {

				bribes++
			}
		}
	}

	fmt.Println(bribes)
}

func main() {
	q := []int32{2, 1, 5, 3, 4}
	minimumBribes(q)
}
