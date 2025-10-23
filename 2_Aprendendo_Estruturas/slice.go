package main

import "fmt"

func main(){
	arr := []string{"Go", "Java", "Python", "JavaScript"}

	slice := arr[0:4]
	
	fmt.Println("Array Original:", arr)
	fmt.Println("Slice:", slice)

	copyExample()
}

func copyExample() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2, slice1)
}
    