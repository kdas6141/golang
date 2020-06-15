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
	"net/http"
	"encoding/json"
	"log"
	"html/template" 
	)

/* individual employee structure */
type Employee struct {
	Name string
	Age uint
	Salary uint 
}

type EmployeeDB struct {

	ifname string 			/* input emplyee database file name */
	ofname string 			/* output emplyee database file name */ 
	employees []Employee
}

var edb EmployeeDB

/* check if employee DB is empty */
func emptyEmployeeDB() bool {
	return len(edb.employees) == 0
}

/* main function of the module */
func main() {
 
	initEmployeeDB()
	if edb.ifname != "" {
		fileRead()
		listEmployees()
		listEmployeesWeb2()
		listEmployeesWeb()
		saveEmployees()
		os.Exit(0);
	}
} /* End of main() */

/* Initialize whole employeeDB system */
func initEmployeeDB() {

	/* initialize input and output file name  */
	edb.ifname = ""
	edb.ofname = ""

	/* get the number of arguments */
	argCount := len(os.Args[1:])
	if argCount == 0  {
		fmt.Print("Missing Employee Input DB File Name...")
		return
	}
	/* save input file name */
	edb.ifname = os.Args[1]
	/* save output file name */
	if argCount == 2 {
		edb.ofname = os.Args[2]
	} else {
		edb.ofname = edb.ifname
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
		e := Employee {Name: eSlice[0]}
		u32, err := strconv.ParseUint(eSlice[1], 10, 32)
    	if err != nil {
        	fmt.Println(err)
    	}
    	e.Age = uint(u32)
		u32_2, err_2 := strconv.ParseUint(eSlice[2], 10, 32)
    	if err != nil {
        	fmt.Println(err_2)
    	}
    	e.Salary = uint(u32_2)
		edb.employees = append(edb.employees, e)
	}
	fmt.Printf("\nFile: %s read done", edb.ifname)
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
		for _, emp := range edb.employees {
			s := fmt.Sprintf("%s;%d;%d\n", emp.Name, emp.Age, emp.Salary)
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
 
type EmployeeOutput struct {
	OneLine string
}

/* List all emplyees and output to stdio using text/format*/
func listEmployees() {
	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB Empty")
	} else {
		var eo EmployeeOutput
 		t := template.New("Employee Record Print")
 		t, _ = t.Parse("{{.OneLine}}")

		fmt.Println("\n# Employee Name           Age             Salary")
		fmt.Println("===================================================")
		for i, emp := range edb.employees {
			eo.OneLine = fmt.Sprintf(" %2d. %-20s %-10d   %10s\n", i+1, emp.Name, emp.Age, int32InsertComma(emp.Salary))
 			t.Execute(os.Stdout, eo)
		}
	}	
} /* End of listEmployees() */

/* Save employee database in output file */
func saveEmployees() {
	if emptyEmployeeDB() {
		fmt.Println("EmployeeDB Empty. Skip file write")
	} else {
		if edb.ofname == "" {
			fmt.Println("Empty output file name")
		}
		if edb.ofname == "" {
			fmt.Println("File name empty. Skip file write")
		} else {
			fileWrite()
		}
	}
}/* End of saveEmployees() */

/* List all emplyees and convert to string */
func employeesDB2Str() string {
	str := ""

	if emptyEmployeeDB() {
		fmt.Println("\nEmployeeDB Empty")
	} else {

		str = fmt.Sprintf("# Employee Name           Age             Salary")
		str += fmt.Sprintf("\n===================================================")
		for i, emp := range edb.employees {
			str += fmt.Sprintf("\n %2d. %-20s %-10d   %10s", i+1, emp.Name, emp.Age, int32InsertComma(emp.Salary))
		}

		str += "\n"
	}
	return str
} /* End of employeesDB2Str() */

/* List employees web handler function based on string */
func printEmployeesWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}     
	fmt.Fprintf(w, employeesDB2Str()) // send data to client side 
}  /* End of printEmployeesWeb() */

/* Enable webserver to display employee records */
func listEmployeesWeb() {
 	http.HandleFunc("/", printEmployeesWeb) 
	// set router     
	err := http.ListenAndServe(":9090", nil)
	// set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} 
}

/* List employees web handler function based on jason */
func printEmployeesWeb2(w http.ResponseWriter, r *http.Request) {
	empInfo, err := json.Marshal(edb.employees)
	if err != nil {
		log.Fatal("Jason Marshal: ", err)
	} 
	fmt.Printf("%s\n", empInfo)

	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}     
	fmt.Fprintf(w, string(empInfo)) // send data to client side 
}  /* End of printEmployeesWeb() */

func listEmployeesWeb2() {
	http.HandleFunc("/", printEmployeesWeb2) 
	// set router     
	err := http.ListenAndServe(":7070", nil)
	// set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} 
}