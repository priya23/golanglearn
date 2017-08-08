package main

import (
"fmt"
"time"
)

func read(re ,re1 chan int){
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

    /*for {
    	 x,ok:= <-re
    		if ok {
            fmt.Printf("Value %d was read.\n", x)
        } else {
            fmt.Println("Channel closed!")
        }
        time.Sleep(time.Second*2)
        continue
      }
    for {
    	 x,ok:= <-re1
    		if ok {
            fmt.Printf("Value %d was read.\n", x)
        } else {
            fmt.Println("Channel closed!")
        }
        time.Sleep(time.Second*2)
        continue
      }*/
      for{_,ok:= <-re
      _,ok1:= <-re1
      if (!ok) {
      	re1 <-1
      	continue 
      }else if (!ok1){
      	re <- 1
      	continue 
      }else {
      	fmt.Println("lock acquired")
      	break
      }
    }
fmt.Println("CAME BACK")
}
func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	ch1 := make(chan int, 2)
	ch1 <- 4
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	go read(ch,ch)
	go read(ch,ch)
	time.Sleep(time.Second * 10)
	
}
