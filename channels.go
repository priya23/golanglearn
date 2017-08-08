package main 
import (
	"fmt"
	"sync"
	"math/rand"
	"github.com/Pallinder/go-randomdata"
	"time"
)
var wg sync.WaitGroup

type Philospher struct{
	name string
	righthand chan int
	lefthand chan int
}


func (ph *Philospher) eating(){
	fmt.Println("started eating")

	for{
		_,ok:= <-ph.righthand
      _,ok1:= <-ph.lefthand
      if (!ok) && (!ok1) {
      	sleeprandom()
      	continue 
      }else if (!ok){
      	ph.lefthand <- 1
      	sleeprandom()
      	continue
      }else if (!ok1){
      	ph.righthand <- 1
      	sleeprandom()
      	continue 
      }else {
      	fmt.Println("lock acquired")
      	break
      }
    }
	fmt.Println("value is read")
}

func (ph *Philospher) stopeatint(){
	ph.righthand <-1
	ph.lefthand <-1
}
func sleeprandom() {
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	return
}

func activity(ph *Philospher){
	defer wg.Done()
	for i:=0;i<1;i++{
		fmt.Println(ph.name + " is hungry")
		ph.eating()
		sleeprandom()
		ph.stopeatint()
		fmt.Println(ph.name + " finished eating")
		fmt.Println(ph.name + " started thinking")
		sleeprandom()
		fmt.Println(ph.name + " finish thinking")
	}
}


func main(){
	var forks [5]chan int
	for i:= range forks{
		forks[i] = make(chan int,1)
		forks[i] <- 1
	}
	var ph = [5]Philospher{}
	for i:=0;i<5;i++{
		ph[i].name=randomdata.FullName(randomdata.Female)
		ph[i].lefthand = forks[i]
		ph[i].righthand = forks[(i+1)%5]
	}

	for j:=0;j<5;j++{
		wg.Add(1)
		go activity(&ph[j])
	}
	wg.Wait()
}