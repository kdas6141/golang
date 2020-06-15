/*
 * Kaushik Das
 * Part B:
 * Following program has a race condition. 
 * Fix the race condition in this code. Rewrite this code
 * in a separate file and use golang channels so there is no race condition.
 * version go1.14.2 windows10
 */
package main

import (
	"fmt"
)

/* global variable to count */
var global int

/* function which performs counting */
func count(finished chan bool) {
	for i := 0; i < 10000; i++ {
		global++
	}
	/* after counting 10000 send finished flag to channel */
	finished <- true
}

func main() {
	/* unbuffered channel */
	finished:=make(chan bool)
	/* call go routine: count first time */
	go count(finished)
	/* received signal from go routine that count has been completed */
	<- finished
	/* call go routine: count second time */
	go count(finished)
	/* received signal from go routine that count has been completed */
	<- finished
	/* print total result */
	fmt.Println("global: ", global)
}
