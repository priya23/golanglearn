package main
import (
 "fmt"
 "strings"
 )


func main() {
	words := strings.Fields("I ate a donut. Then I ate another donut.")
	a := make(map[string]int)
	
	for _, word := range words{
		res,ok := a[word]
		if (ok == true) {
			fmt.Println("inside true")
			fmt.Printf("word is %s",word)
			a[word]=res+1
		} else{
		fmt.Println("not inside true")
			fmt.Printf("word is %s",word)
			a[word]=1
		}
	}
	fmt.Println(a)
}
