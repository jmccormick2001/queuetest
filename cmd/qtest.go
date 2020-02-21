package main

import (
	"fmt"

	"github.com/jmccormick2001/queuetest/queue"
)

func main() {
	q := queue.New()
	q.PushFront("aoo")
	fmt.Printf("q len %d\n", q.Len())
	if q.Len() != 10 {
		q.PushFront("boo")
	}
	fmt.Printf("q len %d\n", q.Len())
	s := q.PopBack()
	fmt.Printf("q len %d\n", q.Len())
	fmt.Printf("got %s\n", s)
}
