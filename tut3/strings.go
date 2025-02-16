package main
import (
	"fmt"
)

func main(){
	var myString string = "resume"
	fmt.Println(myString)
	var indexed =  myString[0]
	fmt.Println(indexed)
	for i, v:= range myString{
		fmt.Println(i , v)
	}

	var strslice = []string{"n"+"i"+"g"}
	var catslice = ""
	for i := range strslice{
		catslice += strslice[i]
	
	}
	println(catslice)

}