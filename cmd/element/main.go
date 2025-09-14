package main
import "fmt"

var buffer [256]byte

func AddOneToEachElement(slice []byte) []byte{
	for i := range slice{
		slice[i]++
	}
	return slice
}

func main(){
	slice := buffer[10:20]
	fmt.Println("raw", slice)
	for i := 0; i<len(slice); i++ { 
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	newSlice := AddOneToEachElement(slice)
	fmt.Println("after original variable", slice)
	fmt.Println("after new variable", newSlice)
}
