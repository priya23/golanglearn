package main 
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)
type Profile struct {
	Name string
	Hobbies []string
}
func main(){
	urls := os.Args[1:]
	argCount := len(urls)
	for i:=0;i< argCount;i++{
		resp ,err := http.Get(urls[i])
		if err != nil{
			fmt.Println(err)
			continue
		}

	fmt.Println("status code:", resp.StatusCode)
	body,err := ioutil.ReadAll(resp.Body)
	var s =new(Profile)
	errr := json.Unmarshal(body,&s)
	if errr != nil{
		fmt.Println(errr.Error())
		continue
	}
	fmt.Println("values are" ,s)
	fmt.Println("****************")
	defer resp.Body.Close()
	}
}