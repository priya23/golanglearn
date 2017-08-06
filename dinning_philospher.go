package main 
import (
	"fmt"
	"time"
)
var ph = []string{"priyanka","karan","sakshi","arun","suresh"}
type fork byte
var hunger = 2
var eat = time.Second /100
var thinking= time.Second /100
func philospher(name string,firsthannd ,secondhand chan fork, complete chan bool){
	//for eating or sleeping
	othertask := func(d time.Duration){
		time.Sleep(100 * d)
	}

	for h := hunger; h > 0; h--  {
		fmt.Println(name+"is hungry")
		<-firsthannd
		<-secondhand
		fmt.Println(name +"done eating")
		othertask(eat)
		firsthannd<- 'f'
		secondhand <- 'f'
		fmt.Println(name+"thinking")
		othertask(thinking)
	}

	complete <-true

}
func main() {
	fmt.Println("table empty")
firsthand := make(chan fork,1)
firsthand <- 'f'
f1:=firsthand
var complete = make(chan bool)
for i:=1; i<len(ph);i++ {
	secondhand := make (chan fork,1)
	secondhand <- 'f'
	go philospher(ph[i],firsthand,secondhand,complete)
	firsthand=secondhand
}
go philospher(ph[0],firsthand,f1,complete)

for range ph{
	<-complete
}

fmt.Println("done")
}