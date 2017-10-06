# Go Problems Worksheet
Hi my name is John Mannion and i'am a thrid year student in GMIT studying software development. As part of a new module called 
data representation and query our class was given the oppourtunity to learn a new programming language called [GO](https://golang.org/)  .
This is a new programming language that has been developed by google and was officially released as Go version 1.0 in March 2012. 

This repository contains a number of go problem solutions. The go problems vary from a simple hello world program to a more complicated 
problem such as newton's method for square roots. The work sheet can be found [here](https://data-representation.github.io/problems/go-fundamentals.html)  .

# Running the code

To run the code in this repository, the files must first be compiled. The Go [compiler](https://golang.org/doc/install) 
must first be installed on your machine. Once that is installed, the code can be compiled and run by following these 
steps. We assume you are using the command line.

## Clone this repository using Git.
1. git clone https://github.com/johnmannion51/go-problems-work-sheet (or other url of git repository)
Change into the folder.
2. cd go-problems-work-sheet
Compile the first file with the following command.
3. go build 01-hello-world.go 
(this will produce an executable file)
Run the executable produced.
4. ./01-hello-world.exe
Hello, world!.
5.Repeat steps 3 and 4 above, replacing the filenames as appropriate.
go compile <filename.go>
./<filename.exe>

# Worksheet
The GO problem worksheet can be found [here](https://data-representation.github.io/problems/go-fundamentals.html)

## List of GO problems
1. Hello World in Japanese
2. Current time and date
3. Fizz Buzz
4. Factorial
5. Guessing Game
6. High low
7. Palindrome
8. Merge sort
9. Newton's square roots
10. Palindrome

## Issues with Go
While using Go I found a frustrating problem when it came to taking in user input. In the guessing game when the program ran
the user would be prompted to input a name and this was immediately followed by an automatic guess. This turned out to be a 
bug within go itself. After a quick google search I knew I wasn't alone when it came to this problem and found that using 
Scan(%guess) as opposed to Scanf("%\n"%guess) resolved my issue. The same outcome could be achieved by using Scanln(%guess).
This appears to be an issue depending on the model of your computer.