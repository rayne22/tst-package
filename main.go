package main

import (
	"fmt"
	"github.com/rayne22/tst-package/ellipse"
)

func main() {
	//Initalise the Init function with value of A,B
	e := ellipse.Init{
		A: 4, B: 2,
	}
	//this will give answer as 0.9749960430435691
	fmt.Println(e.GetEccentricity())
}