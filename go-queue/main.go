package main

import (
	"example/util/myutil"
	"fmt"
)

func main() {

	var queue myutil.ArrayQueue

	queue.Offer(1)
	queue.Offer(2)
	queue.Offer(3)
	queue.Offer(4)
	queue.Offer(5)
	fmt.Println(queue)
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
}
