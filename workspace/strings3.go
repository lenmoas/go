package main

import "fmt"

func main() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	copy(dest, src)

	// for i := range src {
	//		dest[i] = src[i]
	//}

	fmt.Println(dest)
}
