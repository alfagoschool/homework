package main

//Homework 2.3 (Squirrel)

import "fmt"

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	if N*M >= K {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
