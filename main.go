package main

import "fmt"

func main() {
	nomes := []string{"Bruno", "Pedro", "Vinicius", "Eduardo", "Yan"}
	fmt.Println(nomes[:2])
	fmt.Println(nomes[3:])
	rangeOne := nomes[2]
	fmt.Println(rangeOne)
}