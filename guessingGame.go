package main
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
}
//source of code
//http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

//problem with user input
//https://www.socketloop.com/tutorials/golang-number-guessing-game-with-user-input-verification-example
