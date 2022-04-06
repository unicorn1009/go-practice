package main

import (
	"example/util/myutil"
	"fmt"
)

func main() {

	var queue myutil.ArrayQueue

	for i := 0; i < 6; i++ {
		for i := 0; i < 10; i++ {
			queue.Offer(i)
		}

		for i := 0; i < 5; i++ {
			fmt.Println(queue.Poll())
		}
		fmt.Println("=============================")
	}

	fmt.Println(queue)
	fmt.Println(queue.Size())

}
