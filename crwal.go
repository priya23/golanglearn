
package main
import "fmt"

func goroutine1(c chan int){
c <- 1
}

func main(){
c := make(chan int)
//go goroutine1(c)

x:= <-c
fmt.Println(x)
}
