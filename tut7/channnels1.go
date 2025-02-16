package main

import (
	"fmt"
	"math/rand"
	"time"

)
var MAX_CHICKEN_PRICE float32=5
var MAX_tofu float32=6
func main1(){
	var chickenChannel = make(chan string)
	var tofuchannel = make(chan string)
	var websites = []string{"costco","mostco","boscho"}
	for i:= range websites{
		go checkchickenPrices(websites[i],chickenChannel)
		go checkTofuPrices(websites[i],tofuchannel)

	}
	sendMessage(chickenChannel,tofuchannel)
}

func checkchickenPrices(website string ,  chickenChannel chan string){
	for{

		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<MAX_CHICKEN_PRICE{
			chickenChannel<- website
			return
		}

	}
}

func checkTofuPrices(website string, tofuchannel chan string){
	for{
		time.Sleep(time.Second*1)
		var tofuprice = rand.Float32()*20
		if tofuprice<MAX_tofu{
			tofuchannel<- website
			return

		}
	}
}

func sendMessage(chickenChannel ,tofuchannel chan string){
	select{
	case website := <-chickenChannel:
		fmt.Printf("found a deal on the website %v\n", website)

	case website:= <-tofuchannel:
		fmt.Printf("email sent for tofu %v",website)
	}
}

