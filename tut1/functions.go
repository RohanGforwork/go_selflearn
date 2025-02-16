package main
import (
	"fmt"
    "errors"
)
func main(){
 var printValue string = "helloworld"
 printme(printValue)
 num,denom  :=4,2
 var err error
 result ,rem,err:= div(num,denom)
 if err!=nil{
	fmt.Printf(err.Error())
 }else if rem == 0 {
	fmt.Printf("the result is %v ", result)

 }else {
	fmt.Printf("the result is %v and with reminder %v ", result,rem)

 }
 

}

func printme(printValue string){
	fmt.Println(printValue)
}

func div(num int , denom int ) (float64,float64,error){
	var err error 
	if denom ==0 {
		err = errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result float64 = float64(num)/float64(denom)
	rem := float64(num%denom)
	return result,rem,err
}