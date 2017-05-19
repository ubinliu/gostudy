package co

import (
	"fmt"
	"time"
	"strconv"
)

type ProductIds chan string

func productor(name string, productIds ProductIds){
	var id int = 0
	for {
		id = id + 1
		productIds <- name+":"+strconv.Itoa(id)
		fmt.Println(name, " produce id ", id)
		time.Sleep(2 * time.Second)
	}
}

func consumer(name string, productIds ProductIds){
	for {
		id := <-productIds
		fmt.Println(name, " consume: ", id)
	}
}

func GoChannelIntro(){
	//带缓存 10
	productIds := make(ProductIds, 10)
	
	go productor("productor_1", productIds)
	go productor("productor_2", productIds)
	go productor("productor_3", productIds)
	
	go consumer("consumer_1",productIds)
	go consumer("consumer_2",productIds)
	
	var forever chan int = make(chan int)
	
	<-forever
}

