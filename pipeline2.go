package main
import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x:=0; x<100 ; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squarer
	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//Printer in main go routines
	for {
		y, ok2 := <- squares
		if !ok2 {
			break
		}
		fmt.Println(y)
	}
}