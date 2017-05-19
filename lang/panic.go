package lang

import (
	"fmt"
)

func buildPanicCanRecover(){
	
	//匿名函数
	defer func (){
		fmt.Println("buildPanicCanRecover in defer")
		
		if err:= recover(); err != nil{
			fmt.Println(err)
		}
		
		fmt.Println("buildPanicCanRecover out defer")
	}()
	
	fmt.Println("buildPanicCanRecover in build")
	panic("buildPanicCanRecover test panic")
	
	fmt.Println("buildPanicCanRecover never execute 1")
}

func buildPanicCannotRecover(){
	fmt.Println("buildPanicCannotRecover in build")
	panic("buildPanicCannotRecover test panic")
	
	fmt.Println("buildPanicCannotRecover never execute 1")
}

func PanicIntro(){

	//匿名函数
	defer func (){
		fmt.Println("PanicIntro in defer")
		
		if err:= recover(); err != nil{
			fmt.Println(err)
		}
		
		fmt.Println("PanicIntro out defer")
	}()
	
	buildPanicCanRecover()
	
	buildPanicCannotRecover()
	
	fmt.Println("never execute 2")
}



