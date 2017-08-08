package main 
import (
	"fmt"
	"time"
)
//type fork byte
var eat = time.Second /100
func testing(d chan int,i int){
	val := <-d
	fmt.Println("channel value is %v",i  )
	fmt.Println(val)
	//d <- val
}

func main(){

	for i:=0;i<4;i++{
		dd := make(chan int,1)
		dd <- i
		go testing(dd,i)
		fmt.Println(eat)
		time.Sleep(20000)
	}
}
