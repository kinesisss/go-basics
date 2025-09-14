package main
import "fmt"

var buffer [256]byte

func AddOneToEachElement(slice []byte){
	for i := range slice{
		slice[i]++
	}
}

func main(){
	slice := buffer[10:20]
	fmt.Println("raw", slice)
	for i := 0; i<len(slice); i++ { 
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
}