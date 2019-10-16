package main

import(
	"fmt"
)

func main(){

	a := make([]int , 5)
	a = append(a , 4 ,5, 6, 4)
	fmt.Println(a)

	b := make([]int , 0)
	b = append(b , 5, 3, 2, 4)
	fmt.Println(b)

	//下面这一段代码的缺陷
	//对于函数返回值，当是单个返回值时，，如果我们对其不命名，则不用添加括号()，如果我们对单个返回值命名了，就需要添加括号() ;
	// 当是多个返回值时，不论是否给其命名，就必须要添加括号()，注意如果要对其命名，就需要对所有的返回值命名。
	/*func funcMui(x,y int)(sum int,error){
		    return x+y,nil
		}*/

	//make和new的区别
	//new返回的是指针，该指针指向初始化为该类型零值的内存区域。适用于值类型，如数组、结构体等
	//make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
}
