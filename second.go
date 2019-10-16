package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)
	for key,val := range slice {
		//fmt.Printf("value of val : %v\n" , val)
		m[key] = &val
		fmt.Printf("value of val : %v\n" , val)
	}
	for k,v := range m {
		fmt.Println(k,"->",*v)
	}
}
//for-range循环创建每一个元素的副本，也就是说每次循环的时候都是把当前索引位置的值赋值给一个全局变量，而不是元素的引用。
//下面是for-range的循环原理an
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
//从以上的伪代码中可以看到：首先其计算遍历次数（切片的长度），；每次遍历，都会把当前遍历到的值放到一个全局变量value中




