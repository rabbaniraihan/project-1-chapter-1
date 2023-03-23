package helper

import "fmt"

func Help(data []string, c chan<- string) {
	for i := 1; i <= 4; i++ {
		c <- fmt.Sprintln(data, i)
	}
}
