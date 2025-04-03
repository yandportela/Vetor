package main

import "fmt"

func main() {
	inteiros := []int{1, 2, 3, 4, 5}
	inteiros = append(inteiros, 6, 7, 8)
	fmt.Println(inteiros, len(inteiros), cap(inteiros))
}