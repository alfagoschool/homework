package main

import "fmt"

func main() {
    //fmt.Println("hello world")
	
	// Declaring some variables
	var a,b int
	
	fmt.Println("Введите A B:")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	
	// Printing the given texts
    fmt.Printf("A = %d, B = %d, A + B = %d", a, b, a+b)
}