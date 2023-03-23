package main

import (
	"fmt"
	"project-chapter-1/helper"
)

func main() {
	c := make(chan string)
	data1 := []string{"bisa1", "bisa2", "bisa3"}
	data2 := []string{"coba1", "coba2", "coba3"}

	go helper.Help(data1, c)
	go helper.Help(data2, c)

	for i := 0; i < 8; i++ {
		fmt.Println(<-c)
	}

}
