package oo

import (
	"fmt"
	"strconv"
)
//枚举性别
const(
	Unknow = 0
	Male = 1
	Female = 2
)
//定义基类
type Person struct{
	name string //首字母小写，文件外不可见
	age int
	gender int
}

func (p *Person) Init(name string, age int, gender int){
	p.name = name
	p.age = age
	p.gender = gender
}

func (p *Person) Say(words string){
	fmt.Println(words)
}

//定义接口
type Action interface{
	IntroSelf()
}

//定义派生类
type Man struct{
	Person //继承
	// p Person //成员
}
//重写覆盖父类方法
func (m *Man) Init(name string, age int){
	//调用父类的方法
	m.Person.Init(name, age, Male)
}

//继承父类属性和方法，并且实现Action接口
func (m *Man) IntroSelf(){
	m.Say("my name is "+ m.name+", I am "+strconv.Itoa(m.age)+" years old, my gender is "+ strconv.Itoa(m.gender))
}

//定义派生类
type Woman struct{
	Person
}

func (w *Woman) Init(name string, age int){
	w.Person.Init(name, age, Female)
}
//继承父类属性和方法，并且实现Action接口
func (w *Woman) IntroSelf(){
	w.Say("my name is "+ w.name+", I am "+strconv.Itoa(w.age)+" years old, my gender is "+ strconv.Itoa(w.gender))
}
