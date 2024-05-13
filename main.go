package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Welcome to the Test_repository!!")
	var arr [50]string
	fmt.Printf("\n\nCurrent array hosts nothing. Add a few names to your array! It can hold upto %d names!\n\nEnter the number of people you wanna add the names of: ", cap(arr))
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("\nName of #%d person: ", i+1)
		fmt.Scanf("%s", &arr[i])
		if arr[i] == "#stop" {
			break
		} else if arr[i] == "" {
			fmt.Println("Please enter a valid name.")
			i--
		}
	}
	fmt.Printf("\n\nThe names of the participants are as follows:\n")
	for i := 0; i < n; i++ {
		if arr[i] != "#stop" {
			fmt.Printf("%d. %s\n", i+1, arr[i])
		}
	}
}
