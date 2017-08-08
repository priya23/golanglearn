package main 
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
type priya struct{
	name string
	ff *Fork
}
type Fork struct{
	val int
	mux sync.Mutex
}
func acquirelock(p *priya) {
	defer wg.Done()
	p.ff.mux.Lock()
	fmt.Println("lock acquired on ")
}

func main() {
	var f = [5]Fork{}
	
	fmt.Println(f)
	for i:=0;i<5;i++{
		f[i].val=i
	}

	priyanka := priya{name: "priya",ff: &f[1]}
	
	wg.Add(2)
	go acquirelock(&priyanka)
	
	//wg.Add(1)
	go acquirelock(&priyanka)
	
	wg.Wait()

}