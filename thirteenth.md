定义一个包内全局字符串变量，下面语法正确的是（）

- A. var str string
- B. str := ""
- C. str = ""
- D. var str = ""

```
答：A  D
我一开始选的时候，也选择了B选项，关于B选项的:=操作符，它有几个要求点：
1-必须使用显式初始化；2-不能提供数据类型，编译器会自动推导；3-只能在函数内部使用简短形式。
也就是说:=操作符只能定义局部变量。
```



下面这段代码输出什么?

```
func hello(i int) {  
    fmt.Println(i)
}
func main() {  
    i := 5
    defer hello(i)
    i = i + 10
}
```

```
答案： 5
在执行到defer语句的时候，会先把实参计算好，确定函数的实参。
```



下面这段代码输出什么？

```
type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowA()
}
```

```
答案：
showA
showB

结构体嵌套问题
```

















