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
	count := 10
	select {
		case <- time.After(10 * time.Second):
			//do nothing
			fmt.Println(count)
			count = count - 1
		case <- abort:
			fmt.Println("Launch aborted!")
			return
	}
	launch()
}