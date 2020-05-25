/*
 * Kaushik Das
 * Go Program to Implement Parking Lot Simulator
 * version go1.14.2 windows10
 */
package main
import (
	"bufio"
	"os"
	"fmt"
)

const MAX_PARKING_SAPCE = 5

type Node struct {
	val int 				/* Node value */
	next *Node 				/* pointer to next node of the list */
}

type List struct {
	head *Node 				/* pointer to head node of the list */
	tail *Node 				/* pointer to tail node of the list */
	size int                /* size of the list */
}

/* append a node after tail */
func (l *List) PushBack(newNode *Node) {

	l.size++
 	l.tail = newNode
 	if l.head == nil {
 		l.head = newNode
 		return
 	} 
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode 
}  

/* remove a node from tail */
func (l *List) PopBack() *Node{
	var n *Node

	l.size--
 	if l.tail == nil {
 		return nil
 	}

 	n = l.tail
 	if l.tail == l.head {
 		l.head = nil
 		l.tail = nil
 		return n
 	}

 	currentNode := l.head
 	for currentNode.next != l.tail {
 		currentNode = currentNode.next
 	}
 	currentNode.next = nil
 	l.tail = currentNode
 	return n
}  

/* Check if list empty */
func (l *List) Empty() bool {
	return l.size == 0
}

/* get length of the list */
func (l *List) Len() int {
	return l.size
}

/* head node  value */
func (l *List) Front() int {

	if l.size == 0 {
		return -1
	}
	return l.head.val 
}  

/* tail node  value */
func (l *List) Back() int {

	if l.size == 0 {
		return -1
	}
	return l.tail.val 
}  

type Parking struct {

	alleyA *List 							/* parking alley */
	alleyB *List 							/* second alley to retrieve car */
	ticketList [MAX_PARKING_SAPCE]bool 		/* pool of tickets */
	capacity int 							/* maimum capacity of parking structure */
	empty int 								/* total empty slots currently avilable */
}

var park Parking

func main() {
 
	initParking()
	for {
		fmt.Print("\nD) isplay P) ark R) etrieve Q) uit: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _, err := reader.ReadRune()
		if err != nil {
  			fmt.Println(err)
		}
		//windows insert \r\n when user hits enter but linux or macos insert \n
		//to compile in windows i added an extra scanf.
		//fmt.Scanf("%d", &choice)
		switch  choice {
		case 'D', 'd':
			display()
		case 'P', 'p':
			parkCar()
		case 'R', 'r':
			retrieveCar()
		case 'Q', 'q':
			os.Exit(0);
		default:
			fmt.Println("Wrong choice")
		} /* End of switch */
	} /* End of for() */
} /* End of main() */

/* initialize whole parking system */
func initParking() {
	park.alleyA = &List{}
	park.alleyB = &List{}
	park.empty = MAX_PARKING_SAPCE
	park.capacity = MAX_PARKING_SAPCE
	for i:=0; i<MAX_PARKING_SAPCE; i++ {
		park.ticketList[i] = false
	}
}

/* check is particular slot empty */
func emptyParkingSlot(slot int) bool {
	return park.ticketList[slot] == false
}

/* check if there is any empty space in parking space */
func emptyParking() bool {
	return park.empty > 0
}

/* check if the parking space is completely empty */
func completeEmptyParking() bool {
	return park.empty == MAX_PARKING_SAPCE
}

/* acquire an empty space */
func acquireSpace(slot int) {
	park.alleyA.PushBack( &Node{ val: slot+1 } )
	park.ticketList[slot] = true
	park.empty--
}

/* release space of the designated slot */
func releaseSpace(slot int) {
	park.ticketList[slot-1] = false
	park.empty++
	for park.alleyA.Back() != slot {
		park.alleyB.PushBack(park.alleyA.PopBack())
	}
	park.alleyA.PopBack()
	for !park.alleyB.Empty() {
		park.alleyA.PushBack(park.alleyB.PopBack())
	}
}

/* get the parking capacity */
func capacityParking() int {
	return park.capacity
}

/* reserve next parking space */
func reserveNextFreeParkingSlot() {
	if !emptyParking() {
		return
	}

	for i:=0; i<capacityParking(); i++ {
		if  emptyParkingSlot(i) {
			acquireSpace(i)
			fmt.Printf("Ticket no. = %d\n", i+1) 
			break
		}
	} 
}

/* display all acquired slots */
func displayAllSlots() {
	fmt.Print("\nAlley A:")
	for slot := park.alleyA.head; slot != nil; slot = slot.next {
		fmt.Printf(" <%d>", slot.val)
	}
}

/* diplay all cars in the parking space */
func display() {
	if completeEmptyParking() {
		fmt.Println("My Lot is Empty")
	} else {
		displayAllSlots()
	}	
} /* End of display() */

/* park a car */
func parkCar() {
	if !emptyParking() {
		fmt.Println("My Lot is Full")
	} else {
		reserveNextFreeParkingSlot()
	}
} /* End of parkCar() */

/* retrieve a car */
func retrieveCar() {
	if completeEmptyParking() {
		fmt.Println("My Lot is Empty")
	} else {
		var slot int
		fmt.Print("Ticket No.: ")
		fmt.Scanf("%d\n", &slot)
		if slot < 0 || slot > capacityParking() {
			fmt.Printf("Invalid Ticket No: %d", slot)
		} else {
			if emptyParkingSlot(slot-1) {
				fmt.Printf("Already empty slot: %d", slot)
			} else {
				releaseSpace(slot)
			}
		}
	}	
} /* End of retrieveCar() */