package main
 
func fibonacci() func() int {
  first , second , sum:= 0,1,0

  return func() int {
    sum = first 
    first = second
    second = first + second
    //sum ,first,second=first,second,first+second
 return sum
  }

func main(){
  f:=fibonacci()
  for i:= 0 ;i<10 ; i++ {
    fmt.Println(f())
  }
}