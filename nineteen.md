下面代码段输出什么？

```
type Person struct {
    age int
}

func main() {
    person := &Person{28}

    // 1. 
    defer fmt.Println(person.age)

    // 2.
    defer func(p *Person) {
        fmt.Println(p.age)
    }(person)  

    // 3.
    defer func() {
        fmt.Println(person.age)
    }()

    person.age = 29
}
```

```
答案：
29
29
28
首先根据defer先进后出的原则进行分析
第一个输出的肯定是3：
第一个是闭包，会根据上下文语境去引用
第二个输出的是2：
因为person是指针，所以传递是指针，传递的指针的值不变，也就是指向的Person结构体不变，又因为该结构体的字段age由重新赋值为29，所以2输出的是29
最后一个输出的是1：
person.age是函数的参数，是值传递的形式，会把28缓存到栈中，等到最后执行该defer语句的时候取出。
```

