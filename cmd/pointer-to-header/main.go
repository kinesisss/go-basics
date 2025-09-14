package main
import "fmt"

var buffer [256]byte

func SubtractOneFromLength(slicePtr *[]byte){
	slice := *slicePtr // de-referencing the pointer to access the original object
	*slicePtr = slice[0:len(slice)-1] // the assignation must be done to the original object, thus the de-referencing
}

func main(){
	slice := buffer[100:120]
	fmt.Println("slice length before subtracting", len(slice))
	SubtractOneFromLength(&slice)
	fmt.Println("slice length after subtracting", len(slice))
}

