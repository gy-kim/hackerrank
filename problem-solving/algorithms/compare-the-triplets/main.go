package main

import "fmt"

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	res := []int32{0, 0}

	for i := 0; i < len(a); i++ {
		switch {
		case a[i] > b[i]:
			res[0]++
		case a[i] < b[i]:
			res[1]++
		}
	}

	return res
}

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}

	fmt.Println(compareTriplets(a, b))
}
