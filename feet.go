package main

import "fmt"

func main(){
	var input float64
	const mul = 0.3048
	fmt.Print("Feet :")
	fmt.Scanf("%f",&input)
	fmt.Printf("\n%f ft = %f m\n",input,mul*input)
}
