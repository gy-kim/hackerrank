package main

import "fmt"

// https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	cnt := int32(0)
	socks := map[int32]int{}

	for i := int32(0); i < n; i++ {
		no := ar[i]
		_, ok := socks[no]
		if ok {
			cnt++
			delete(socks, no)
		} else {
			socks[no] = 1
		}
	}

	return cnt
}

func main() {
	n := int32(9)
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(n, ar))
}
