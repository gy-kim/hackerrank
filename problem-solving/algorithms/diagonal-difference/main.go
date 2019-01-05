package main

import (
	"fmt"
)

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {
	var res, sum1, sum2 int32 = 0, 0, 0
	size := len(arr[0])

	for i := 0; i < len(arr[0]); i++ {
		sum1 += arr[i][i]
		sum2 += arr[i][size-i-1]
	}

	res = sum1 - sum2
	if res < 0 {
		res *= -1
	}

	return res
}

func main() {
	arr := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	fmt.Println(diagonalDifference(arr))
}
