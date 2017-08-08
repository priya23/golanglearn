package main 
import (
	"fmt"
	"sync"
	"time"
)

func priya(k *sync.Mutex){
	k.Lock()
	//fmt.Println("before firstlocked")
	//k.Lock()
	fmt.Println("locked")
	time.Sleep(200000)
	fmt.Println("came out")
	//k.Unlock()
	return
}
func main(){
	lh1:= &sync.Mutex{}
	f1:=lh1
	fmt.Printf("%p",lh1)
	fmt.Printf("%p",f1)
	go priya(lh1)
	time.Sleep(2000000)
	fmt.Printf("%p",lh1)
	fmt.Printf("%p",f1)
	//fmt.Println("came out")
	//time.Sleep(200000)
	//go priya(lh1)
	//for i:=1; i<4;i++ {
	//	rh:= &sync.Mutex{}
	//	go priya(rh)
	//}
	time.Sleep(time.Minute * 2)
}