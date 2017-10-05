/*package main
import (
    "fmt"
    "math/rand"
    "time"
)

func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := xrand(1, 10)
    tries := 0
    var guess int
	var prev int

    fmt.Println("Welcome to Guess My Number Game!")
    for guess != myrand {
        fmt.Println("Take a guess...")
        fmt.Scanf("%d\n", &guess)
        
		if guess == prev {
		fmt.Println("Try again you already tried this number")
		tries--
		}
		tries++
        if guess > myrand {
            fmt.Println("Guess too high")
        } else if guess < myrand {
            fmt.Println("Guess too low")
        } else {
            fmt.Printf("Good job! You guessed it in %v tries", tries)
            break
        }
		
		prev = guess
    }
}*/

package main
//imports required for program
import(
    "fmt"
    "math/rand"
    "time"
)//import

//this function will be called upon to generate random numbers
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}//xrand

func main() {
    var myname string
	//calls xrand function to generate a random number between 1 - 6
    myrand := xrand(1, 6)
	//declared variables for user input
    guessTaken := 0
    var guess int
	//prompts the user to enter a name and explains the rules of the game
    fmt.Println("Hello! What is your name?")
    fmt.Scanf("%s", &myname)
    fmt.Printf("Hello, %s, I am thinking of a number between 1 and 6.\n", myname)
    
    //this is the for loop used to prompt the user to take a guess 
	//if the users guess is too low they are prompted to have another go
	//if the users guess is too hign they are prompted to have another go

    for guessTaken < 6 {
        fmt.Println("Have a guess...")
        fmt.Scan(&guess)
        guessTaken++
        if guess < myrand {
            fmt.Println("Your guess is too low.")
        }
        if guess > myrand {
            fmt.Println("Your guess is too high.")
        }
        if guess == myrand {
            break
        }
    }
	//informs the user if they have guessed correctly or 
	//lets them know what the number was if they guessed incorrectly
    if guess == myrand {
        fmt.Printf("Good job %s! You guessed my number was %d in %d tries\n", myname, myrand, guessTaken)
    } else {
        fmt.Printf("Hard luck %s. The number I had in mind was %d\n", myname, myrand)
    }
}//main
//source of code
//http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

//http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html

//problem with user input
//https://www.socketloop.com/tutorials/golang-number-guessing-game-with-user-input-verification-example
