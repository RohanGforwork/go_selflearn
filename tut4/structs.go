package main
import(
	"fmt"

)
type gasEngine struct{
	mpg uint8 
	gallons uint8
	
}
type electric struct{
	mpkw uint8
	battery uint8
}

type allengine interface{
	milesleft() uint8
}

func main(){
	var myEngine gasEngine = gasEngine{8,2}
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg,myEngine.gallons)

	fmt.Println(myEngine.milesleft())

	var myEngine1 electric = electric{2,4}

	fmt.Println(myEngine1.lengthleft())
	canmakeit(myEngine1,50)



}

func (e gasEngine) milesleft() uint8{
	var milesLeft  = e.gallons*e.mpg
	return milesLeft

}

func (e electric) milesleft() uint8{
	var milesLeft = e.mpkw*e.battery
	return milesLeft
}


func (e electric) lengthleft() uint8{
	var lengthLeft  = e.mpkw*e.battery
	return lengthLeft
}


func canmakeit(e gasEngine,miles uint8) {
	if miles <= e.gallons*e.mpg{
		fmt.Println("can make it")

	}else{
		fmt.Println("cannot make it ")
	}
}