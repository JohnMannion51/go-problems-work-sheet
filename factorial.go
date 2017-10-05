//Program to find Factorial of number
package main
import
	"fmt"
/* Variable Declaration */
var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
                       // Range: 0 through 18446744073709551615. 
var i int = 1
var n int
 
/*     function declaration        */
func factorial(n int) uint64 {   
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
            factVal *= uint64(i)  // mismatched types int64 and int
        }
         
    }    
    return factVal  /* return from function*/
}
 
func main(){    
    fmt.Print("Enter a positive integer between 0 - 50 : ")
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",factorial(n))
}

/*
package main

import (
	"fmt"      // Import for input/output
	"math/big" // Import for math
	"strconv"  // import to convert string
)

func digitSum(userinput int) int {

	factorial := big.NewInt(1) 

	// calculate factorial
	for i := 1; i <= userinput; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i))) 
	}
	fmt.Println(userinput,"!  : ",factorial)//output factorial

	var factSum, digit int  

	//calculate sum of factorial digits
	for i := range factorial.String() {
			digit, _ = strconv.Atoi(string(factorial.String()[i])) 
			factSum += digit                                 
	}
	return factSum
	
}//digitSum

func main() {
	//take user input from console
	var userinput int
	fmt.Print("Please enter a number(1-100): " )
	fmt.Scan(&userinput)

	fmt.Println("Input : ",userinput)
	fmt.Println("Sum of Factorial Digits : ",digitSum(userinput))
}*/