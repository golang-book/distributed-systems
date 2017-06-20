package lvl1

import "fmt"

func Print(in <-chan string) {
	for msg := range in {
		fmt.Println(msg)
	}
}
