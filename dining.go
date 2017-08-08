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
	righthand *Fork
	lefthand *Fork
}

func sleeprandom() {
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	return
}

func(ph *Philospher) eating(){
	fmt.Println(ph.name ,"trying to get lock")
	ph.righthand.mux.Lock()
	time.Sleep(time.Second * 1 )
	ph.lefthand.mux.Lock()

	// critical area.

	fmt.Printf("right %s %p", ph.name, &ph.righthand.mux)
	fmt.Printf("left %s %p", ph.name, &ph.lefthand.mux)
	fmt.Println(ph.name, "acquired lock")
}

func(ph *Philospher) stopeatint(){
	ph.righthand.mux.Unlock()
	ph.lefthand.mux.Unlock()
	fmt.Println(ph.name + " released lock")
}
type Fork struct{
	val int
	mux sync.Mutex
}

func activity(ph *Philospher){
	defer wg.Done()
	for i:=0;i<10;i++{
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

	//create 5 forks
	var f = [5]Fork{}

	//create philosphers
	var ph = [5]Philospher{}
	for i:=0;i<5;i++{
		f[i].val=i
		ph[i].name=randomdata.FullName(randomdata.Female)
		ph[i].lefthand = &f[i]
		ph[i].righthand = &f[(i+1)%5]
	}

	for j:=0;j<5;j++{
		wg.Add(1)
		go activity(&ph[j])
	}
	//wg.Add(1)
	//go activity(&ph[4])
	wg.Wait()

}