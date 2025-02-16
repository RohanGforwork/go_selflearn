package main

import(
	"fmt"
	"sync"
)
var newString = []float64{1,2 ,3,4,5}
var wg = sync.WaitGroup{}
var m = sync.Mutex{}
func main(){
wg.Add(1)//add the constrain to WaitGroup
go dbcall(newString)

wg.Wait()//wait until wg.done
}

func dbcall(newString []float64) {

	for i := range newString{
		fmt.Printf("\nthe number present is %v \n", newString[i])
		m.Lock()//deadlock condition
		newAppend(newString[i]) 
		m.Unlock()//same 
		wg.Done()//complete the thing
	}
	
}

func newAppend(interger float64){
	fmt.Printf("\nthe value of the new square is %v\n",interger*interger)
}