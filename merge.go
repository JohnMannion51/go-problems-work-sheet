package main

import (
	"fmt"	//import for input and output
	"sort"	//import for sorting
)//import
//function for the merge sort
func mergeSort() {
		//creates 2 arrays of int type unsorted
		a := []int{8, 12, 22}
		b := []int{11, 6, 50}
		//returns new slice and then this merged slice is passed as separate argument
		ab := append(a, b...)
	
		//sorts and merges the arrays
		sort.Ints(ab)
		//prints both unsorted arrays and merged array
		fmt.Println("List A: ",a)
		fmt.Println("List B: ",b)
		fmt.Println("List A and list B Merged and Sorted:",ab)
}//meregSort

func main(){
//calls the mergeSort function for use
	mergeSort()
}//main

/*sources 
https://stackoverflow.com/questions/45076887/what-does-append-do-in-go
https://gobyexample.com/arrays

*/