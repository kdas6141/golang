/*
 * Kaushik Das
 * version go1.14.2 windows10
 */
/* MAGIC SQUARE - An NxN matrix containing values from 1 to N*N that are  */
/* ordered so that the sum of the rows, columns, and the major diagonals  */
/* are all equal.  There is a particular algorithm for odd integers, and  */
/* this program constructs that matrix, up to 13 rows and columns.  This  */
/* program also adds the sums of the row, columns, and major diagonals.   */

package main
import "fmt"

func main() {
	var (
   		input int                               /* User defined integer       	*/
   		loc[][] int                         	/* Array holding all          	*/
      	otherdiag int                          	/* Total of one matrix diagonal */
    )
	/*                                                                        	*/
	/* Print introduction of what this program is all about.                  	*/
	/*                                                                        	*/
   fmt.Println("\nMagic Squares: This program produces an NxN matrix where")
   fmt.Println("N is some positive odd integer.  The matrix contains the")
   fmt.Println("values 1 to N*N.  The sum of each row, each column, and")
   fmt.Println("the two main diagonals are all equal.  Not only does this")
   fmt.Println("program produces the matrix, it also computes the totals for")
   fmt.Println("each row, column, and two main diagonals.")

   fmt.Println("\nBecause of display constraints, this program can work with")
   fmt.Println("values up to 13 only.")

   fmt.Println("\nEnter a positive, odd integer (-1 to exit program):");

   for {
		//windows insert \r\n when user hits enter but linux or macos insert \n
		//to compile in windows i added an extra scanf.

   		fmt.Scanf("%d\n", &input)
		/*                                                                        */
		/*    If input = -1, then exit program.                                   */
		/*                                                                        */
      	if input == -1 {
	 		break
      	}
		/*                                                                        */
		/*    Validity check for input: Must be a positive odd integer < 13.      */
		/*                                                                        */
      	if input <= 0 {
         	fmt.Println("Sorry, but the integer has to be positive.")
	 		fmt.Println("Enter a positive, odd integer (-1 to exit program):")
  	 		continue
      	}
      	if input > 13 {
	 		fmt.Println("Sorry, but the integer has to be less than 15.")
	 		fmt.Println("Enter a positive, odd integer (-1 to exit program):")
	 		continue
      	}
      	if input%2 == 0 {
	 		fmt.Println("Sorry, but the integer has to be odd.")
	 		fmt.Println("Enter a positive, odd integer (-1 to exit program):");
	 		continue;
      	}

      	loc = createMatrix(input)
      	initMatrix(input, loc)
      	otherdiag = 0
     	calculateMatrix(input, loc, &otherdiag)
     	displayMatrix(input, loc, otherdiag)

      	fmt.Println("Enter a positive, odd integer (-1 to exit program):");

	}   /* End of for input>-1 loop */
	fmt.Println("\nBye bye!");
} /* End of main() */

// Use make to create two dimensional array
// allocate memory for 1-dimensional array of size matrixSize
// within for loop allocate memory for 2-dimension array
func createMatrix(matrixSize int) (twoDim [][]int) {
	 twoDim = make([][]int, matrixSize+1)

	 for row := 0; row <= matrixSize; row++ {
	 	twoDim[row] = make([]int, matrixSize+1)
	 }
	 return twoDim 
}

// Create nested for loop to initialize array
func initMatrix(matrixSize int, loc [][]int) {
	/*                                                                      */
	/*    Initialize Matrix 												*/
	/*                                                                      */
    for row := 0; row <= matrixSize; row++ {   	/* Initialize matrix with 	*/
        for col := 0; col <= matrixSize; col++ {/* all zeroes.              */
            loc[row][col] = 0
        } 
    }                                
}

// Convert C code calculation without changing the logic.
func calculateMatrix(matrixSize int, loc [][]int, otherdiag *int) {

	/* Values will reside within  	*/
	/* rows 1 to input*input and  	*/
    /* columns 1 to input*input.  	*/
    /* Row totals will reside in  	*/
    /* loc[row][0], where row is  	*/
    /* the row number, while the  	*/
    /* column totals will reside  	*/
    /* in loc[0][col], where col  	*/
    /* is the column number.      	*/

	/* Initialize  row, col 		*/
    row := 1   								/* First value gets to sit on */
    col := matrixSize/2 + 1                 /* 1st row, middle of matrix. */

 	/*                                                                        */
	/*    Loop for every value up to input*input, and position value in matrix*/
	/*                                                                        */
    for value := 1; value <= matrixSize*matrixSize; value++ { /* Loop for all values.    */
        if loc[row][col] > 0 {            		/* If some value already     */
        		                                /* present, then             */
           	row += 2                    		/* move down 1 row of prev.  */
            if row > matrixSize {              	/* If exceeds side, then     */    
               	row -= matrixSize               /* go to other side.         */
            }

            col--                          		/* move left 1 column.       */
            if col < 1 {                	    /* If exceeds side, then     */
               	col = matrixSize   	        	/* go to other side.         */
            }
        }

        loc[row][col] = value       	      	/* Assign value to location.  */

		/*                                                                    */
		/*       Add to totals                                                */ 
		/*                                                                    */
        loc[0][col] += value              		/* Add to its column total.   */
        loc[row][0] += value 	             	/* Add to its row total.      */
        if (row == col) {                    	/* Add to diagonal total if   */
            loc[0][0] += value             		/* it falls on the diagonal.  */
        }

        if (row+col == matrixSize+1) {	        /* Add to other diagonal if   */
            *otherdiag += value;     	        /*  it falls on the line.     */
        }

		/*                                                                    */
		/*       Determine where new row and col are                          */
		/*                                                                    */
        row--
        if row < 1 {		                    /* If row exceeds side then   */
        	row = matrixSize	                /*  goto other side.          */
        }
        col++
        if col > matrixSize {		            /* If col exceeds side then   */
            col = 1                        		/*  goto other side.          */
      	}                                     	/* End of getting all values. */
   	}

}

//Create nested for loop to read each array element and display
//matrix with column totals and row totals
func displayMatrix(matrixSize int, loc [][]int, otherdiag int) {

	/*                                                                         */
	/*    Print out the matrix with its totals                                 */
	/*                                                                         */
    fmt.Print("\nThe number you selected was ", matrixSize)
    fmt.Println(", and the matrix is:\n")
    for row := 1; row <=matrixSize; row++ {     /* Loop: print a row at a time */
        fmt.Print("     ");               	 	/* Create column for diag.total*/
        for col := 1; col <=matrixSize; col++ {
        	fmt.Printf("%5d", loc[row][col]) 	/* Print values found in a row */
        }
        fmt.Printf(" = %5d\n",loc[row][0]);  	/* Print total of row.         */
    }

 	/*                                                                        */
	/*    Print out the totals for each column, starting with diagonal total. */
	/*                                                                        */
    for col := 0; col <=matrixSize; col++ {     /* Print line separating the  */
        fmt.Print("-----")                		/* value matrix and col totals*/
    }
    fmt.Printf("\n%5d",otherdiag);          	/* Print out the diagonal total*/
    for col := 1; col <=matrixSize; col++ {
	    fmt.Printf("%5d",loc[0][col])      		/* Print out the column totals*/
    }
    fmt.Printf("   %5d\n",loc[0][0])       		/* Print out the other diagonal*/
                                           		/*  total                      */
} 