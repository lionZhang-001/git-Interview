package main

import (
	"fmt"
)

func main(){

	sn1 := struct {
		age int
		name string
	}{age : 11 , name : "zyy"}

	sn2 := struct {
		age int
		name string
	}{age : 11 , name : "zyy"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age : 12 , m : map[string]string{"a" : "1"}}

	sm2 := struct {
		age int
		m map[string]string
	}{age : 12 , m : map[string]string{"a" : "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
//结构体比较的问题
//1-结构体只能比较是否相等，但不能比较大小；2-相同类型的结构体才能进行比较，结构体是否相等不但与属性类型相关，还与属性顺序相关。
//3-若struct内的所有成员可以比较，则该struct可以通过==或！=进行比较是否相等。比较时逐项进行比较，如果每一项都相等，两个结构体才相等，否则不相等。
//可以进行比较的类型：bool  数值型  字符  指针 数组；而切片  map  函数都是不可比较的
