package main

import "fmt"

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {

	return append(a[d:], a[0:d]...)

}

func main() {
	a := []int32{1, 2, 3, 4, 5}
	d := int32(4)
	fmt.Println(rotLeft(a, d))
}
