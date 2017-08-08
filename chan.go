package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

//for fork as chan
var fork = make(chan int, 5)

//string of philosphers
var ph = []string{"priyanka", "karan", "sakshi", "arun", "suresh"}


//cycle to eat for randon time and think
func consume(ph string,wg *sync.WaitGroup) {
	defer wg.Done()
	//function to sleep for randon time
	randomsleep := func() {
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	}
	for {
		fmt.Println("consming" + ph)
		<-fork
		<-fork
		randomsleep()
		fmt.Println("finish eating" + ph)
		fork <- 1
		fork <- 1
		//planning to think
		fmt.Println(ph +"thinking")
		randomsleep()
	}
	fmt.Println(ph + "completed")
	//ch <- true
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		fork <- 1
	}

	//channel to keep track if philopher is finished eating
	//var status = make(chan bool)

	for i := 0; i < len(ph); i++ {
		wg.Add(1)
		go consume(ph[i],&wg)
	}
	wg.Wait()
	/*for k := 0; k < len(ph); k++ {
		<-status
	}*/
}
