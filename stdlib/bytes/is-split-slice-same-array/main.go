package main

import (
	"bytes"
	"fmt"
)

func main() {

	arr := [10]byte{1, 2, 73, 4, 5, 73, 7, 8, 73, 10}
	fmt.Println("arr: ", arr)

	sl := arr[:]
	fmt.Println("sl: ", sl)

	slsl := bytes.Split(sl, []byte{73})
	fmt.Println("slsl:")
	fmt.Println("arr: ", arr)
	fmt.Println("sl: ", sl)
	fmt.Println("slsl: ", slsl)

	for _, el := range slsl {
		el[0] = el[0] + 100
	}
	fmt.Println("slsl:")
	fmt.Println("arr: ", arr)
	fmt.Println("sl: ", sl)
	fmt.Println("slsl: ", slsl)
}
