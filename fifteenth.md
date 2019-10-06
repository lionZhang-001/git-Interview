下面代码下划线处可以填入哪个选项？

```
func main() {
    var s1 []int
    var s2 = []int{}
    if __ == nil {
        fmt.Println("yes nil")
    }else{
        fmt.Println("no nil")
    }
}
```

- A. s1
- B. s2
- C. s1、s2 都可以

```
答案：C
s1是nil切片，表示一个不存在的切片；s2是空切片

go版本1.13，根据官方的解释，nil is a predeclared identifier representing the zero value for a
 pointer, channel, func, interface, map, or slice type.
 也就是说nil是这些类型的零值。
 关于切片的三种状态：https://www.jianshu.com/p/eccc2aa3d0bd
```



下面这段代码输出什么？

```
func main() {  
    i := 65
    fmt.Println(string(i))
}
```

- A. A
- B. 65
- C. compilation error

```
答案：A
在go语言中，字符串的内部实现是通过utf-8编码实现的。可以通过rune类型方便的对每一个utf-8字符进行访问。
字符串中的每一个元素叫做字符，go语言中的字符有以下两种：
一种是uint8类型,或者叫做byte型，代表了ascii码的一个字符；
另一种是rune类型，代表一个utf-8字符，在处理中文、日文或其他符合字符时，就需要用到rune类型，rune类型实际上是一个int32类型。
```



下面这段代码输出什么？

```
type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowA())
    fmt.Println(b.ShowB())
}
```

```
答案：
13 
23
结构体是值类型，赋值是通过值拷贝的形式来完成的，也就是说a b c三者的地址都不同。
另外如果想打印结构体c，需要通过fmt.Printf("%p" , &c)的形式格式化输出。


```

