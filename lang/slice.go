package lang

import (
	"fmt"
)

func SliceIntro(){
	//初始化一个slice变量，初始长度0，容量10
	var nums = make([]int, 0, 10)
	
	fmt.Println("len:", len(nums), "cap:", cap(nums))
	
	//插入
	nums = append(nums, 1,2,3,4,5,6,7,8,9,10,11)
	
	fmt.Println("after append change len:", len(nums), "cap:", cap(nums))
	
	//截取片段
	fmt.Println(nums[6:10])
	
	//非引用传值
	NoChangeSliceValue(nums)
	
	fmt.Println("after nochange len:", len(nums), "cap:", cap(nums))
	//引用传值
	ChangeSliceValue(&nums)
	
	fmt.Println("after change len:", len(nums), "cap:", cap(nums))
	//遍历
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
	}
}

func ChangeSliceValue(s *[]int){
	*s = append(*s, 12, 13, 14,15,16,17,18,19,20)
}

func NoChangeSliceValue(s []int){
	s = append(s, 12, 13, 14,15,16,17,18,19,20)
}