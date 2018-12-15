package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	jumps := int32(0)

	for i := 0; i < len(c)-1; i++ {
		if len(c) > i+2 && c[i+2] == 0 {
			i++
		}
		jumps++
	}
	return jumps
}

func main() {
	c := []int32{0, 0, 1, 0, 0, 0, 0}

	fmt.Println(jumpingOnClouds(c))
}
