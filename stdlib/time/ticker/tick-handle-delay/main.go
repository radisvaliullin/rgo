package main

import (
	"fmt"
	"time"
)

func main() {

	tk := time.NewTicker(time.Second)

	for i := 0; i < 10; i++ {
		t := <-tk.C
		fmt.Println("tick:", t)
		if i == 5 {
			fmt.Println("sleep")
			time.Sleep(time.Second * 5)
		}
	}
}
