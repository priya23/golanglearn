package main

import (
"fmt"
"time"
)

func read(re chan int){
	//v1,ok1 := <-re
	//fmt.Println(v1)
	//fmt.Println(ok1)
	fmt.Println("st start")
	/*fmt.Println("st start")
	select {
	case x,ok:= <-re: 
		if ok {
            fmt.Printf("Value %d was read.\n", x)
        } else {
            fmt.Println("Channel closed!")
        }
    default:
        time.Sleep(time.Second*2)
        continue
    }*/

    for {
    	 x,ok:= <-re
    		if ok {
            fmt.Printf("Value %d was read.\n", x)
        } else {
            fmt.Println("Channel closed!")
        }
        time.Sleep(time.Second*2)
        continue
      }
    

}
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	go read(ch)
	go read(ch)
	go read(ch)
	time.Sleep(time.Second * 10)
	
}
