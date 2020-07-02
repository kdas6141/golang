package main
import (
	"fmt"
	"time"
	"os"
)

func launch() {
	fmt.Println("Vehicle has been launched!")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown>0; countdown-- {
		fmt.Println(countdown)
		select {
		case <- tick:
			//do nothing
		case <- abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}