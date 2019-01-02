package main

import "fmt"

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	swap := int32(0)
	for i := int32(0); i < int32(len(arr)); i++ {
		if (i + 1) != arr[i] {
			t := i
			for arr[t] != i+1 {
				t++
			}
			temp := arr[t]
			arr[t] = arr[i]
			arr[i] = temp
			swap++
		}
	}
	return swap
}

func main() {
	arr := []int32{2, 3, 4, 1, 5}

	fmt.Println(minimumSwaps(arr))
}
