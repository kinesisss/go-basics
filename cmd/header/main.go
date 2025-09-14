package main

import "fmt"

var buffer [256]byte 

func SubtractOneFromLength(slice []byte) []byte{
	slice = slice[0:len(slice)-1]
	return slice
}

func main(){
	slice := buffer[100:150]
	fmt.Println("length of slice before: ", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("length of slice (original array) after subtraction: ", len(slice))
	fmt.Println("length of newSlice that should contain the modification: ", len(newSlice))
}
