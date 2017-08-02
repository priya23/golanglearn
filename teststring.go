package main
import (
 "fmt"
 //"strings"
 )


func main() {
	var vaal = "hello my name is samantha"
	for i , word := range vaal{
		fmt.Printf("%v is %q" ,i,word)
		fmt.Println("******")
	}
}
