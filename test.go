package main

import "fmt"

func fruit() {
	fruits := []string{"plumb", "apple", "orange", "pear"}

	fmt.Println(fruits[0]); // plumb

	fmt.Println(fruits[0:2]) // [plumb apple]
	fmt.Println(fruits[:2])  // [plumb apple]
	fmt.Println(fruits[1:2])  // [apple]
	fmt.Println(fruits[2:])  // [orange pear]
}
