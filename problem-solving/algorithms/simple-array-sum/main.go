package main

import "fmt"

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for _, a := range ar {
		sum += a
	}
	return sum
}

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}
	fmt.Println(simpleArraySum(ar))
}
