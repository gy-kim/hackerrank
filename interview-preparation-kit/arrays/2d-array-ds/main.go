package main

import (
	"fmt"
	"math"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	max := int32(math.MinInt32)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := int32(0)
			sum += arr[i][j] + arr[i][j+1] + arr[i][j+2]
			sum += arr[i+1][j+1]
			sum += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func main() {
	arr := [][]int32{}

	a1 := []int32{1, 1, 1, 0, 0, 0}
	a2 := []int32{0, 1, 0, 0, 0, 0}
	a3 := []int32{1, 1, 1, 0, 0, 0}
	a4 := []int32{0, 0, 2, 4, 4, 0}
	a5 := []int32{0, 0, 0, 2, 0, 0}
	a6 := []int32{0, 0, 1, 2, 4, 0}

	arr = append(arr, a1, a2, a3, a4, a5, a6)

	fmt.Println("arr", arr)

	fmt.Println(hourglassSum(arr))
}
