package main

import (
	"fmt"
	. "math"
)

func pow_Δ(l float64) float64{
	if l == 0.0{
		return 0.0
	}else if l < 0{
		return -1.0
	}else{
		return Pow(l,0.5)
	}
}

func main(){
	var (
		a float64
		b float64
		c float64
	)

	fmt.Println("a:")
	fmt.Scanln(&a)
	fmt.Println("b:")
	fmt.Scanln(&b)
	fmt.Println("c:")
	fmt.Scanln(&c)

	fmt.Println(a,b,c)

	var (
		Δ float64
		x1 float64
		x2 float64
	)

	Δ = b*b - 4*a*c
	var lin_shi float64 = pow_Δ(Δ)
	fmt.Println(Δ,lin_shi)

	if (lin_shi < 0){
		fmt.Println("无解")
	}else{
		x1 = (-b + lin_shi)/(2*a)
		x2 = (-b - lin_shi)/(2*a)
		fmt.Println("x1=",x1," ", "x2=",x2)
	}

	main()
}