package main

import(
	"fmt"
)

func main(){
    //问题在于append第一个参数要求是切片类型，而不是指针类型
	list := new([]int)
	fmt.Printf("%T \n" , list) //返回的是指向切片的指针
	list = append(list, 1) //append第一个参数要求是切片
	fmt.Println(list)

	//append(s2 , s1...)或者就是append(s2 , 1 ,2 ,3)
	s1 := []int{ 1, 3, 4}
	s2 := []int{ 3, 5, 2}
	s2 := append(s2 , s1)

	//形如size := 1024这种声明形式，是有限制的：1-必须使用显式初始化；2-不能提供数据类型，编译器会自动推导；3-只能在函数内部使用简短形式。
	//在这里var ( size = 1024)  这种声明方式是对的
	var (
		size := 1024
		max_size := size * 2
	)
	fmt.Println(size , max_size)

}
