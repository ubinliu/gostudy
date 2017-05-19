package co

import (
	"fmt"
	"time"
)

func setTimeout(seconds int, timeout chan bool){
	time.Sleep(time.Duration(seconds)*time.Second)
	timeout <- true
}

func sum(result chan int){
	time.Sleep(time.Duration(30)*time.Second)
	ret := 0
	for i:= 0; i < 1000; i++ {
		ret = ret + i
	}
	result<-ret
}

func TimeoutIntro(){

	timeout := make(chan bool)
	result := make(chan int)
	//设置超时时间
	go setTimeout(10, timeout)
	//抛出协程
	go sum(result)
	
	select {
		case s := <-result:
			fmt.Println("sum result:", s)
		case <-timeout:
			fmt.Println("timeout")	
	}
	
	fmt.Println("will exit")
}