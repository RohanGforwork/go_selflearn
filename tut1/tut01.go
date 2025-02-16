package main 
import "fmt"
func tut1(){
	var intNum int64 = 19 
	fmt.Println(intNum)

	var floatNum float64 = 34.3
	fmt.Println(floatNum)

	var result float64 = floatNum * float64(intNum)
	fmt.Println(result)

	var myString string = "hello pp!!"+ ""+ "world"
	fmt.Println(myString)

	fmt.Println((len(myString)))

	var myboolean bool = false
	fmt.Println(myboolean)

	const mycosnt string = "const value"
	fmt.Println(mycosnt)
}