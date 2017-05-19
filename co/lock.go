package co

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

var counter int = 0

func Incr(name string, locker *sync.Mutex){
	for{
		locker.Lock()
		
		counter++
		fmt.Println(name," incre:", counter)
		
		locker.Unlock()
		
		time.Sleep(1*time.Second)
	}
}

func LockerIntro(){
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	locker := &sync.Mutex{}
	
	go Incr("counter1", locker)
	go Incr("counter2", locker)
	go Incr("counter3", locker)
	
	var forever chan int
	
	<-forever
	
}

