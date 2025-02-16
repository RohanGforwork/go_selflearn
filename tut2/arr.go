package main

import(
	"fmt"
)

func main(){
	var intArr [3]int32
	intArr[1]= 123
    fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	var newArr []int32 = []int32{1,2,3}
	newArr = append(newArr,7)
	fmt.Println(newArr)

	var intSlice2 []int32 = []int32{1,2,3,4,4}
	newArr = append(newArr, intSlice2...)

	var intslice3 []int32 = make([]int32, 3,12) //length,capacity
    fmt.Println(intslice3) 


	//mapping

	var myMap map[string]uint8  = make(map[string]uint8)
	fmt.Println(myMap)
    var myMap2 = map[string]uint8{"rohan":20,"eren":23}
	fmt.Println(myMap2["rohan"])
	fmt.Println(&intArr[0])

	var isAge , ok = myMap2["jason"]
	if ok{
		fmt.Printf("is present %v",isAge) //ok=true

	}else{
		fmt.Printf("not prisent")
	}
    delete(myMap2,"eren")

	for name:= range myMap2{
		fmt.Printf("name = %v /t/t age = %v",name ,isAge)


	}
	///for i,v :=range intArr{
	///}
	///for {
	///if i>=10 {
	///	i++
	///	break
		
	///}
	///}

	///for i:=0 ; i<=10;i++{
		
	///}
}