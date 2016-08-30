package main

import "fmt"

func main(){
	var C,F float64
	fmt.Print("Température en F :")
	fmt.Scanf("%f",&F)
	C = (F-32)*5/9
	fmt.Printf("%f °F = %f °C\n",F,C)
}
