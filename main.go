package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Welcome to the Test_repository!!")
	var n int
	fmt.Printf("\n\nEnter the size of the pattern (for example, 3 for 3x3): ")
	fmt.Scanf("%d", &n)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	fmt.Printf("\n")
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
