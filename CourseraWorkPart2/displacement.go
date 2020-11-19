package main

import (
	"fmt"
	"math"
)
var(
	acce float64
	initV float64
	initD float64
)
func main(){

	fmt.Printf("Please Enter Acceleration: ")
	_, err := fmt.Scanf("%f", &acce)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Please Enter Init Velocity: ")
	_, err = fmt.Scanf("%f", &initV)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Please Enter Init Displacement: ")
	_, err = fmt.Scanf("%f", &initD)
	if err != nil {
		fmt.Println(err)
	}

 fn := GenDisplaceFn(acce,initV,initD)
 fmt.Println(fn(3))
}

func GenDisplaceFn (a  float64, initVel float64,initDis float64)func(float64)float64{
	fn:= func (time float64)float64{
		time2 := math.Pow(time,2.0)
		a = a*time2
		initVel = initVel*time
		s := (0.5*a)+initVel+initDis
		return s
	}
	return fn
}