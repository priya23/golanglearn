package main
import (
	"net/http"
	"fmt"
	"os"
	)


func main(){

	urls := os.Args[1:]
	argCount := len(urls)
	for i:=0;i< argCount;i++{
		resp ,err := http.Get(urls[i])
		if err != nil{
			fmt.Println(err)
		 	continue
			//fmt.Println("error is ",err)
		}
	fmt.Println(urls[i])
	fmt.Println("response is " ,resp)
	fmt.Printf("***************\n")
	defer resp.Body.Close()
	}
	
}