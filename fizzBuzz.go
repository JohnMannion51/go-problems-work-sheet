package main
//imports fmt package
import "fmt"
//use for loop to run from 1 - 100
func main() {
    for i := 1; i <= 100; i++ {
	
        if i%3 == 0 {
            // Multiples of 3 are viewed as fizz
            fmt.Printf("fizz")
        }
        if i%5 == 0 {
            // Multiples of 5 are viewed as buzz
            fmt.Printf("buzz")
        }

        if i%3 != 0 && i%5 != 0 {
            // Multiples of both 3 and 5 are viewed as fizzbuzz
            fmt.Printf("%d", i)
        }

        // A trailing new line (so both fizz + buzz can be printed on the same line)
        fmt.Printf("\n")

    }
}
//source of code
//https://golangcode.com/fizz-buzz-test-in-go/