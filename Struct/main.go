package main

import (
	"fmt"
)

func main() {
	num := 100
	update(&num)
	fmt.Println(num)
}

func update(num *int) {
	*num = 10
}
