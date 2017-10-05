//Program to find Factorial of number
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
	//take in user input
	var userinput int
	fmt.Print("Please enter a number(1-100): " )
	fmt.Scan(&userinput)

	fmt.Println("Input : ",userinput)
	fmt.Println("Sum of factorial digits is ",digitSum(userinput))
}