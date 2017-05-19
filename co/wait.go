package co

import (
	"fmt"
	"sync"
	"runtime"
	"time"
	"strconv"
)

var s int = 0
var wg sync.WaitGroup

func Wait(name string){
	s++
	t := s
	time.Sleep(time.Duration(t)*time.Second)
	fmt.Println(name," sleep:", t)
	defer wg.Done()
}

func WaitIntro(){
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Wait("process_"+strconv.Itoa(i))
	}
	
	wg.Wait()
	
	fmt.Println("process done")
}

