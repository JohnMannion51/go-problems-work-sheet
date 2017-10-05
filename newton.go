package main

import "fmt"//import for input/output
import "math"// import for math (math.Sqrt(x))


	//function to run newtons method
	func newton(x float64,z float64)float64{
		//var y float64  =	z - (((z*x) - x) / (2*z))
		return z - (((z*z) - x) / (2*z))
	
	}//newton

	func main(){
		
		var counter int =0;
		var x float64;
		//number to calculate square root
		//take user input from console
		fmt.Printf("Newtons Theory Square Root Calculator")
		fmt.Printf("\n=====================================")
		fmt.Print("\nPlease enter a number to test:")
		fmt.Scan(&x)
		fmt.Print("\n")

		//initial guess
		z :=  float64(1)

		//loop through using newtons method until next guess == current guess
		for z = 1.0; z != newton(x,z); z = newton(x,z){
		//for z = 1.0; z_next(x,z)<math.Sqrt(x); z = z_next(x,z){
			//display current guess
			counter++
			fmt.Printf("Guess %d :%12.10f \n",counter,z)
		}//for

		fmt.Printf("The square of %.2f is %12.10f according to Newtons Theory\n",x,z)
		fmt.Printf("The square of %.2f is %.8f according to 'math.Sqrt(x)'",x,math.Sqrt(x))

	}//main