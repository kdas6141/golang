package main  // Example ex8-08 
 
import (
	"fmt"
	"strconv"
	"bytes"
//	"html/template"  
	"net/http"
	"encoding/json" 
) 
 
/* individual employee structure */
type Employee struct {
	Name string
	Age uint
	Salary uint 
}
 
type EmployeePageData struct {  
	PageTitle 	string
	PageBar 	string  
	EmpInfo		[]Employee
} 

var epd EmployeePageData

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

func main() {
 
	//tmpl := template.Must(template.ParseFiles("layout_employee.html")) 
 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        json.NewDecoder(r.Body).Decode(&epd.EmpInfo)
        fmt.Println("Inside decoder")
     	fmt.Println("\n# Employee Name           Age             Salary")
		fmt.Println("===================================================")
		for i, emp := range epd.EmpInfo {
			fmt.Printf("\n %2d. %-20s %-10d   %10s", i+1, emp.Name, emp.Age, int32InsertComma(emp.Salary))
		}

/*        fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)

    	data := TodoPageData {
        	PageTitle: "My TODO list",
        	Todos: []EmpInfo {
        		{ Title: "Task 1", Done: false},     
        		{ Title: "Task 2", Done: true},     
        		{ Title: "Task 3", Done: true},    
        	},   
    	}   

    	tmpl.Execute(w, data)
*/
   	}) 

	http.ListenAndServe(":7070", nil) 
}