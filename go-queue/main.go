package main

import (
	"example/util/myutil"
	"fmt"
)

func main() {

	var queue myutil.ArrayQueue

	for i := 0; i < 14; i++ {
		queue.Offer(i)
	}
	fmt.Println(queue)
	fmt.Println(queue.Size())
	fmt.Println("=============================")
	for i := 0; i < 4; i++ {
		queue.Poll()
	}
	queue.Offer(14)
	fmt.Println(queue)
	fmt.Println(queue.Size())
	fmt.Println("=============================")

	queue.Offer(15)
	queue.Offer(16)
	queue.Offer(17)
	queue.Offer(18)
	fmt.Println(queue)
	fmt.Println(queue.Size())
	fmt.Println("=============================")

	queue.Offer(19)
	queue.Offer(20)
	fmt.Println(queue)
	fmt.Println(queue.Size())

	// var a []int
	// a = make([]int, 5)
	// fmt.Println(a, len(a), cap(a))
	// fmt.Printf("%p\n", &a)
	// a = append(a, 3)
	// fmt.Println(a, len(a), cap(a))
	// fmt.Printf("%p\n", &a)

}
