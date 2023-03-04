# numgo
Fundamental package for array computing in Golang

NumGo is the fundamental package for scientific computing with Golang.

It provides:

a powerful N-dimensional array object


///////////////////////////////////////////////////////////////

go get github.com/nesazeri/numgo

go get -u github.com/nesazeri/numgo


package main

import (

	"fmt"
	"github.com/nesazeri/numgo"
  
)

func main(){

	arrData := []float64{1, 2, 3, 4, 5}
	arr := numgo.Array(arrData)
	fmt.Println("Array: ", arr)
  
}

go run .\test.go
