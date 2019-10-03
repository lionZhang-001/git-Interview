关于channel，下面语法正确的是()

- A. var ch chan int
- B. ch := make(chan int)
- C. <- ch
- D. ch <-

```
答： A  B  C 
我在自己作答的时候选上了D选项，记住：向channel中写入数据需要带上写入的值
A、B都是声明;C是读取channel，可以不用将读取的值赋值给其他变量的;
```



下面这段代码输出什么？

- A.0
- B.1
- C.Compilation error  

```
type person struct {  
    name string
}

func main() {  
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
```

```
答案：A
打印一个map中不存在的值，返回元素类型的零值。这个例子中m的类型是map[person]int，因为map中不存在任何元素，更不存在p，所以打印int类型的零值
```



下面这段代码输出什么？

- A.18
- B.5
- C.Compilation error  

```
func hello(num ...int) {  
    num[0] = 18
}

func main() {  
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
```

```
答案： A
可变参数：如果函数的参数是可变参数，同时还有其他的参数。可变参数要放在形参列表的最后；一个函数的参数列表中，最多只能有一个可变参数。
```















