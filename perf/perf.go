package perf

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"runtime/pprof"
	"os"
	"os/signal"
	"syscall"
)

func compute(){
	sum := 0
	for i := 0; i < 1000000; i++{
		sum = sum + i
	}
}

func httpGet(url string) (result string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get "+url+" failed", err)
		return ""
	}
 
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read fail", err)
	}
	return string(body)
}
var resps []string
func api(){
	resps = make([]string, 0)
	for{
		compute()
		resps = append(resps, httpGet("http://www.163.com"))
		compute()
		//fmt.Println("request:",len(resps))
		time.Sleep(time.Second)
	}
}

func signalprocess(){
	c := make(chan os.Signal)
	//USR1 USR2
	//kill -s USR1 PID
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2)
	var cpufile *os.File
	var memfile *os.File
	
	for{
		sig := <-c
		switch sig {
			case syscall.SIGUSR1:
				cpufile, err := os.OpenFile("./cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
			    if err != nil {
			        fmt.Println("open cpu.prof failed")
			        break
			    }
			    pprof.StartCPUProfile(cpufile)
			    
			    memfile, err := os.OpenFile("./mem.prof", os.O_RDWR|os.O_CREATE, 0644)
				if err != nil {
					cpufile.Close()
				    fmt.Println("open mem.prof failed")
			        break
				}
				pprof.WriteHeapProfile(memfile)
			    
				break
			case syscall.SIGUSR2:
				pprof.StopCPUProfile()
				cpufile.Close()
				memfile.Close()
		}
	}
}
//go tool pprof study ./mem.prof
//go tool pprof study ./cpu.prof
func PprofIntro(){
	
	go signalprocess()
	
	for i := 0; i < 100; i++ {
		go api()
	}
	
	var forever chan int = make(chan int, 1)
	
	<-forever
	
}