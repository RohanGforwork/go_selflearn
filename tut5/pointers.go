package main

import (
	"fmt"
)

func main(){
	var newString = [5]float64{1,2,3,4,5}
	var p *int32 = new((int32))
	var i int32 = 10
	fmt.Println(p)
	*p =10 
	p =&i

	fmt.Printf("the value of i is %v and value of pointer p is %v \n", p, i)
    maniString(&newString) // no need for new variable since func calling and no return function
	fmt.Println(newString)
}

func maniString(secstring *[5]float64) {
	for i := range secstring {
		(*secstring)[i] = (*secstring)[i] * (*secstring)[i] //explicit pointers for accesing local  array newString
	}
}