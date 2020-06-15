/*
 * Kaushik Das
 * Part A:
 * Following program has a race condition. 
 * Fix the race condition in this code. Now rewrite this code using golang channels.
 * version go1.14.2 windows10
 */
package main

import (
	"fmt"
)

func main() {
	/* create an unbuffered channel */
	ch := make(chan int)

	go func(ch chan int) {
		for i := 1; i <= 5; i++ {
			/* write value i into channel */ 
			ch <- i
		}
		/* close the channel */
		close(ch)
	}(ch)

	/* receive value d from channel until channel has been closed */
	for d := range ch {
		fmt.Println("i: ", d)
	}
}