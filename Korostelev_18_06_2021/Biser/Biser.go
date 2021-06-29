package main

import "fmt"

func main() {	
	// Declaring some variables
	var n int
	
	fmt.Println("Введите количество цветов бусин N:")
	fmt.Scanf("%d", &n)

	// Printing the given texts
    fmt.Printf("Количество цветов бусин N = %d.\nМин.число бусин, которые можно вытащить не глядятак, чтобы среди них гарантированно были две бусины одного цвета = %d", n, n+1)
}