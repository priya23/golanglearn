package main
import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	)

func main(){
	urls := os.Args[1:]
	argCount := len(urls)
	for i:=0;i< argCount;i++{
		resp ,err := http.Get(urls[i])
		if err != nil{
			panic(err)
		}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("body is" ,body)
	defer resp.Body.Close()
	}
	
}