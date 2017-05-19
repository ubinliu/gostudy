package oo

import (
)

//多态
func SomeBodyIntroSelf(action Action){
	action.IntroSelf()
}

func ObjectOrientedIntro(){
	//name 和 age都是私有变量，不可见
	//wrong := &Man{name:"name", age:1}
	
	xiaoming := &Man{}
	xiaoming.Init("xiaoming", 13)
	
	xiaohua := &Woman{}
	xiaohua.Init("xiaohua", 12)
	
	SomeBodyIntroSelf(xiaohua)
	
	SomeBodyIntroSelf(xiaoming)
}

