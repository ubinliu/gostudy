package lang

import (
	"fmt"
)

type kv  map[string]string

func MapIntro(){
	//初始化
	cache := make(kv, 3)
	
	fmt.Println("init cache, len:", len(cache))
	//赋值
	cache["a"] = "123"
	cache["b"] = "456"
	
	fmt.Println("init cache, len:", len(cache))
	
	//遍历
	for k, v := range(cache) {
		fmt.Println("k:",k, ",v:", v)
	}
	
	//删除
	delete(cache, "a")
	
	//取值
	v, ok := cache["a"]
	
	if ok {
		fmt.Println("exist a:", v)
	}else{
		fmt.Println("a is deleted")
	}
}

