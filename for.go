package main


import "fmt"

func main(){
  fmt.Println(`1
2
3
4
5`)
  fmt.Println("**********")
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i++
  }
  fmt.Println("**********")
  for i:=1; i<= 10; i++ {
    fmt.Println(i)
  }
}
