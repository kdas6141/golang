/*
 * Kaushik Das
 * Go Program to Implement Employee Database
 * version go1.14.2 windows10
 */
package main
import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"bytes"
	"runtime"
	)

/* user choice */
const ( 
		ADD = '1' + iota 
		DELETE
		SEARCH 
		LIST
		SAVE
		EXIT
	)

/* individual employee structure */
type Employee struct {
	name string
	age uint
	salary uint 
}

/* individual node structure */
type Node struct {
	emp *Employee 			/* Node content: pointer to employee structure */
	next *Node 				/* pointer to next node of the list */
}

/* Linked list of node structure */
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

/* Insert a node into a sorted list */
func (l *List) InsertSorted(newNode *Node) {

	l.size++
	/* empty list */
 	if l.head == nil {
 		l.head = newNode
 	 	l.tail = newNode
 		return
 	} 

 	/* find the position to insert */
	currentNode := l.head
	var prev *Node
	for currentNode != nil  && strings.Compare(newNode.emp.name, currentNode.emp.name) == 1 {
		prev = currentNode
		currentNode = currentNode.next
	}

	/* if position located before head */
	if currentNode == l.head {
		newNode.next = l.head
		l.head = newNode
	} else if currentNode == nil { /* if position located after tail */
		prev.next = newNode
		l.tail = newNode
	} else { /* all other cases */
		prev.next = newNode
		newNode.next = currentNode
	}
}  

/* Delete a node from the list */
func (l *List) DeleteList(name string) {

	l.size--
	/* single item in list */
 	if l.head == l.tail {
 		l.head = nil
 	 	l.tail = nil
 		return
 	} 

 	/* Find the entry to delete */
	currentNode := l.head
	var prev *Node = nil
	for currentNode != nil  && 
		strings.Compare(strings.ToUpper(name), strings.ToUpper(currentNode.emp.name)) != 0 {
		prev = currentNode
		currentNode = currentNode.next
	}

	/* If entry not found */
	if currentNode == nil {
		fmt.Println("Node not found for name: %s", name)
		l.size++
		return
	}

	/*If the last entry to be removed */
	if (currentNode == l.tail) {
		prev.next = nil;
		l.tail = prev
	}

	/* if the first entry to be removed */
	if (currentNode == l.head) {
		l.head = currentNode.next
		currentNode.next = nil
	} else { /* middle entry to be removed */
		prev.next = currentNode.next
	}
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

/* head node value */
func (l *List) Front() *Node {

	if l.size == 0 {
		return nil
	}
	return l.head 
}  

/* tail node value */
func (l *List) Back() *Node {

	if l.size == 0 {
		return nil
	}
	return l.tail
}  

type EmployeeDB struct {

	ifname string 			/* input emplyee database file name */
	ofname string 			/* output emplyee database file name */ 
	elist List 				/* list of employees */
	emap map[string]*Node 	/* pointer to node for employee */
}

var edb EmployeeDB

/* check if employee DB is empty */
func emptyEmployeeDB() bool {
	return edb.elist.size == 0
}

/* main function of the module */
func main() {
 
	initEmployeeDB()
	if edb.ifname != "" {
		fileRead()
	}
	userChoice()	
} /* End of main() */

/* Get user choice from the keyboard */
func userChoice() {
	for {
		fmt.Println("\nMenu Options:")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Delete Employee")
		fmt.Println("3. Search Employee")
		fmt.Println("4. List All Employees")
		fmt.Println("5. Save Employee Database")
		fmt.Println("6. Exit Employee Database")
		fmt.Print("Enter Your Choice: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _, err := reader.ReadRune()
		if err != nil {
  			fmt.Println(err)
		}
		switch  choice {
		case ADD:
			addEmployee()
		case DELETE:
			deleteEmployee()
		case SEARCH:
			searchEmployee()
		case LIST:
			listEmployees()
		case SAVE:
			saveEmployees()
		case EXIT:
			saveEmployees()
			os.Exit(0);
		default:
			fmt.Println("Wrong choice")
		} /* End of switch */
	} /* End of for() */
} /* End pf userChoice */

/* Initialize whole employeeDB system */
func initEmployeeDB() {

	/* initialize linked list and hash function */
	edb.elist.size = 0;
	edb.elist.head = nil
	edb.elist.tail = nil
	if edb.emap == nil {
		edb.emap = make(map[string]*Node)
	}
	edb.ifname = ""
	edb.ofname = ""

	/* get the number of arguments */
	argCount := len(os.Args[1:])
	if argCount == 0  {
		fmt.Print("Missing Employee Input DB File name...")
		return
	}
	/* save input file name */
	edb.ifname = os.Args[1]
	/* save output file name */
	if argCount == 2 {
		edb.ofname = os.Args[2]
	} else {
		edb.ofname = ""
	}
} /* End of initialize employee DB */


/* Read from file and store employee information to the employeeDB */
func fileRead() {
	
	fileHandle, err := os.Open(edb.ifname) 	/* Open the file */
    if err != nil { 
    	fmt.Println(err)
    	os.Exit(1)
    } 

	defer fileHandle.Close()
	/* Create a new Scanner for the file */
	fileScanner := bufio.NewScanner(fileHandle)
	/* read one line at a time from input file */
	for fileScanner.Scan() {

		eSlice := make([]string, 3) 
		eSlice = strings.Split(fileScanner.Text(), ";")
		e := &Employee {name: eSlice[0]}
		u32, err := strconv.ParseUint(eSlice[1], 10, 32)
    	if err != nil {
        	fmt.Println(err)
    	}
    	e.age = uint(u32)
		u32_2, err_2 := strconv.ParseUint(eSlice[2], 10, 32)
    	if err != nil {
        	fmt.Println(err_2)
    	}
    	e.salary = uint(u32_2)
		n := &Node{emp: e, next: nil}
		edb.elist.InsertSorted(n)
		edb.emap[strings.ToUpper(eSlice[0])] = n
	}
	fmt.Printf("\nFile: %s raed done", edb.ifname)
}/* End of file read */


/* Write to file of all employee information from the employeeDB */
func fileWrite() {
	
	fileHandle, err := os.Create(edb.ofname) 	/* Create the file */
    if err != nil { 
    	fmt.Println(err)
    	os.Exit(1)
    } 

	defer fileHandle.Close()

	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB is Empty")
	} else {
		/* retrieve one employee at a time and write into file */
		for n := edb.elist.head; n != nil; n = n.next {
			s := fmt.Sprintf("%s;%d;%d\n", n.emp.name, n.emp.age, n.emp.salary)
			l, err := fileHandle.Write([]byte(s))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if l != len(s) {
				fmt.Println("failed to write data")
				os.Exit(1)
			}		
		}
	}
	fmt.Printf("\nFile: %s write done", edb.ofname)
}/* End of file write */

/* convert a unsigned integer to comma seperated string */ 
func int32InsertComma(val uint) string {
	s := strconv.Itoa(int(val))
	n := 3
    var buffer bytes.Buffer
    for i, rune := range s {
    	if i != 0 && (len(s)-i) % n == 0 {
    		buffer.WriteRune(',')	
    	} 
        buffer.WriteRune(rune)
    }

    return buffer.String()
}

/* List all emplyees in the list */
func listEmployees() {
	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB Empty")
	} else {
		fmt.Println("# Employee Name           Age             Salary")
		fmt.Println("===================================================")
		for i, n := 1, edb.elist.head; n != nil; i, n = i+1, n.next {
			fmt.Printf("\n %2d. %-20s %-10d   %10s", i, n.emp.name, n.emp.age, int32InsertComma(n.emp.salary))
		}
		fmt.Printf("\n");
	}	
} /* End of listEmployees() */

/* find an employee from hash function */
func findEmployee(name string) *Node {

	np, ok:= edb.emap[strings.ToUpper(name)]
	if ok == true { /* found key */
		return np
	}

	return nil
}

/* Read user input string from console */
func userInputString(prompt string) string {

	fmt.Printf("\n%s", prompt)
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
  		fmt.Println(err)
  		return ""
  	}
  	if runtime.GOOS == "windows" {
  		name = strings.TrimRight(name, "\r\n") 		/* for windows */
  	} else {
		name = strings.TrimRight(name, "\n") 		/* for linux */
	}
	return name
}

/* Read user input unsigned int from console */
func userInputUint(prompt string) uint {

	fmt.Printf("\n%s", prompt)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
  		fmt.Println(err)
  		return 0
  	}
  	if runtime.GOOS == "windows" {
  		str = strings.TrimRight(str, "\r\n") 		/* for windows */
  	} else {
		str = strings.TrimRight(str, "\n") 		/* for linux */
	}
	if str == "" {
		return 0
	}
	u64, err_2 := strconv.ParseUint(str, 10, 32)
	if err_2 != nil {
		fmt.Println(err_2)
	}

	return uint(u64)
}

/* Add an employee in Employee DB*/
func addEmployee() {
	for {
		name := userInputString("ADD Enter Employee Name: ")
		if name == "" {
			break
		}
  		np := findEmployee(name)
  		if np != nil {
  			fmt.Printf("\nExisting Employee: %s found", name)
  			break	
  		} else {
  			age:= userInputUint("Add Enter Employee Age: ")
  			if age==0 {
  				fmt.Printf("\nMissing Employee: %s Age", name)
  				break	  				
  			} else {
  				salary:= userInputUint("Add Enter Employee Salary: ")
	  			if salary==0 {
  					fmt.Printf("\nMissing Employee: %s Salary", name)
  					break	  				
  				} else {
					e := &Employee {name: name, age: age, salary: salary}
					n := &Node{emp: e, next: nil}
					edb.elist.InsertSorted(n)
					edb.emap[strings.ToUpper(name)] = n
  					fmt.Printf("\nEmployee: %s has been added", name)
  				}
  			}
  		}
  	}	
} /* End of addEmployee() */

/* Delete an employee  from the Employee DB*/
func deleteEmployee() {
	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB Empty")
	} else {
		for {
			name := userInputString("DELETE Enter Employee Name: ") 
  			np := findEmployee(name)
  			if np == nil {
  				fmt.Printf("\nEmployee: %s not found", name)
  				break	
  			} else {
  				edb.elist.DeleteList(name)
  				delete(edb.emap, strings.ToUpper(name))
  				fmt.Printf("\nEmployee: %s has been deleted", name)
  			}
  		}
	}	
} /* End of deleteEmployee() */

/* Search an employee in Employee DB*/
func searchEmployee() {
	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB Empty. Nothing to search")
	} else {
		for {
			name := userInputString("SEARCH Enter Employee Name: ") 
  			np := findEmployee(name)
  			if np == nil {
  				fmt.Printf("\nEmployee: %s not found", name)
  				break	
  			} else {
  				fmt.Printf("\nEmployee Details:")
  				fmt.Printf("\nName: %s", np.emp.name)
  				fmt.Printf("\nAge: %d", np.emp.age)
  				fmt.Printf("\nSalary: %s", int32InsertComma(np.emp.salary))
  			}
		}
	}	
} /* End of searchEmployee() */

/* Save employee database in output file */
func saveEmployees() {
	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB Empty. Skip file write")
	} else {
		if edb.ofname == "" {
			edb.ofname = userInputString("SAVE Enter File Name: ")
		}
		if edb.ofname == "" {
			fmt.Println("File name empty. Skip file write")
		} else {
			fileWrite()
		}
	}
} /* End of saveEmployees() */