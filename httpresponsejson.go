package main 
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
type Profile struct {
	Name string
	Hobbies []string
}
func main(){
	url := "http://localhost:3000"
	res,err := http.Get(url)
	if err != nil{
		panic(err.Error())
	}
	body,err := ioutil.ReadAll(res.Body)
	var s =new(Profile)
	errr := json.Unmarshal(body,&s)
	if errr != nil{
		panic(errr.Error())
	}
	fmt.Println(s)
}