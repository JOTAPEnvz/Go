package main

import "fmt"

func main() {
	var array1 = [5]int{1, 2, 3, 4, 5}
	var array2 = [5]int{1, 2, 2, 4, 5}
	var array3 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array1 == array2)
	fmt.Println(array1 == array3)
	fmt.Println(array2 == array3)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}