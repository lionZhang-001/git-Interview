
1. 通过指针变量p访问其成员变量name，有哪几种方式

  A p.name  B (&p).name  C (*p).name   D p->name
```bash
答案： A C
```
2. 下面这段代码能否通过编译？为什么？若通过，输出什么内容
```bash
package main
 
import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
    var i int =0
    var i1 MyInt1 = i 
    var i2 MyInt2 = i
    fmt.Println(i1,i2)
}
```
```bash
答案： 类型别名和类型定义的区别
type MyInt int语句是基于int类型定义了一个新的类型，而type MyInt2 = int是给int类型起了一个别名。而go语言是强类型转换，它不允许类似于c++一样进行隐式转换，它只允许显式类型转换。
```
3. 以下代码输出什么？
```bash
func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)

}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}
```
```bash
答案： append()函数是copy了原切片的一个副本，也就是说ap()函数里的切片a和外面的切片a不是一个切片。
所以输出是：
 7 8 9
 7 8 9 
 1 8 9
```

