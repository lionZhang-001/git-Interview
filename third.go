package main

import (
	"fmt"
)

func main(){

	s :=  make([]int , 5)
	fmt.Printf("地址：%v\n" , &s[0])
	s = append(s , 1 , 2 , 3)
	fmt.Printf("地址：%v\n" , &s[0])
	fmt.Println(s)

	s2 := make([]int , 0)
	s2 = append(s2 , 1, 3, 5)
	fmt.Println(s2)
}
