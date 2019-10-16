1.下面代码中，x 已声明，y 没有声明，判断每条语句的对错。

```
1. x, _ := f()
2. x, _ = f()
3. x, y := f()
4. x, y = f()
```

```
答案：
1 错
2 对
3 对
4 错
第一道题我一开始是做错的，这里_就是指求丢弃了该值，所以第一题里面只有一个变量x，而x已经声明。
```



下面代码输出什么？

```
func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}

func increaseB() (r int) {
    defer func() {
        r++
    }()
    return r
}

func main() {
    fmt.Println(increaseA())
    fmt.Println(increaseB())
}
```

- A. 1 1
- B. 0 1
- C. 1 0
- D. 0 0

```
答案：B
重点是理解return这句话
```



下面代码输出什么？

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
    var a A = Work{3}
    s := a.(Work)
    fmt.Println(s.ShowA())
    fmt.Println(s.ShowB())
}
```

- A. 13 23
- B. compilation error

```
答案：A
类型断言，a.(work)得到的是底层类型。
```

