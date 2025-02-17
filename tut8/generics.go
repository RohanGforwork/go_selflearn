package main
import "fmt"

type gasEngine struct{
	mpg int32
	gallons int32
}

type electric struct{
	mpkw int32
	charge int32
}

type car [T gasEngine | electric] struct{
	name string
	model string
	engine T 
}

func main(){
	var newCar = car[gasEngine]{
		name : "lambo",
		model : "a",
		engine: gasEngine{
			mpg: 21,
			gallons: 32,

		},



	}
	fmt.Printf("the name of the cr is %s and the model is %s and the milage given is %v",newCar.model,newCar.name,(newCar.engine.mpg*newCar.engine.gallons)/2)
}