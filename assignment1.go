/*
 * Kaushik Das
 * Go Program to Implement a Queue using Array
 * version go1.14.2 windows10
 */
package main
import (
	"fmt"
	"os"
)

const MAX = 50 + iota

type Choice int32
const (
	Insert Choice = 1 + iota
	Delete
	Display
	Quit
)

var (
	queue_array [MAX] int32
	rear int32 = -1
	front int32 = -1
)

func main() {

	var choice Choice
	for {
		fmt.Println("\n1.Insert element to queue")
		fmt.Println("2.Delete element from queue")
		fmt.Println("3.Display all elements of queue")
		fmt.Println("4.Quit")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		//windows insert \r\n when user hits enter but linux or macos insert \n
		//to compile in windows i added an extra scanf.
		//fmt.Scanf("%d", &choice)
		switch  choice {
		case Insert:
			insert()
		case Delete:
			delete()
		case Display:
			display()
		case Quit:
			os.Exit(1)
		default:
			fmt.Println("Wrong choice")
		} /* End of switch */
	} /* End of for() */
} /* End of main() */

func insert() {
	var add_item int32
	if rear == MAX -1 {
		fmt.Println("Queue overflow")
	} else  {
		if front == -1 {
			/*If queue is initially empty */
			front = 0
		}
		fmt.Print("\nInsert the element in queue: ")
		fmt.Scanf("%d", &add_item)
		//fmt.Scanf("%d", &add_item)
		rear = rear + 1
		queue_array[rear] = add_item 
	}
} /* End of insert() */

func delete() {
	if front == -1 || front > rear {
		fmt.Println("Queue Underflow")
		return
	} else  {
		fmt.Println("Element deleted from queue is: ", queue_array[front])
		front = front + 1
	}
} /* End of delete() */

func display() {
	var i int32
	if front == -1 {
		fmt.Println("Queue is empty")
	} else  {
		fmt.Println("Queue is :")
		for i=front; i<=rear; i++ {
			fmt.Println(queue_array[i])
		}
		fmt.Println("")
	}
} /* End of display() */
