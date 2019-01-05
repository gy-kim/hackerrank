package main

import "fmt"

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	res := int64(0)

	for _, a := range ar {
		res += a
	}

	return res
}

func main() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(ar))
}
