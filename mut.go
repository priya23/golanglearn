package main 
import (
	"fmt"
	"time"
	"sync"
	"math"
	"reflect"
)
type Fork struct{
  v int
  mux sync.Mutex
}
var f = []Fork{
	Fork {v: 1},
	Fork {v: 2},
	Fork {v: 3},
	Fork {v: 4},
	Fork {v: 5},
}
func occupy(i int){
	fmt.Println("inside occupy")
	fmt.Println(f[i].v)
}
func main(){
	
	for i:=0;i<5;i++{
		//fmt.Println(f[i].v)
		//fmt.Println((i%5)+1)
		var  rem = math.Mod(float64(i),5)
		var rem1=int(rem+1)
		fmt.Println(reflect.TypeOf(rem))
		fmt.Println(f[rem1].v)
		//go occupy(i)
		//fmt.Println(f.mux)
	}
time.Sleep(200000)
}
