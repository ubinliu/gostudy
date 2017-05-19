package lang

import (
	"fmt"
)

type Person struct{
	Age int
	Name string
}

func StructIntro(){
	//初始化方式
	p1 := &Person{18, "zhangsan"}
	//返回引用
	p2 := &Person{Name:"zhaosi", Age:19}
	//返回值
	p3 := Person{4, "child"}
	//返回值
	p4 := Person{}
	//返回引用
	p5 := new(Person)
	
	fmt.Println(p1, p2, p3, p4, p5)
	
	NoChangeStruct(p3)
	
	fmt.Println("after no change:", p3)
	
	ChangeStruct(&p3)
	
	fmt.Println("after change:", p3)
	
	
}

func NoChangeStruct(p Person){
	p.Age = 25
}

func ChangeStruct(p *Person){
	p.Age = 25
}