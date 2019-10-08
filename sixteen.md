切片 a、b、c 的长度和容量分别是多少？

```
func main() {

    s := [3]int{1, 2, 3}
    a := s[:0]
    b := s[:2]
    c := s[1:2:cap(s)]
}
```

```
答案：a的长度是0，容量是3；b的长度是2，容量是3；c的长度是1，容量是2。
截取数组的操作[i:j]或[i:j:k]，当i/j省略的时候，默认是0/底层数组的长度，截取到的切片长度和容量分别是j-i，k-i。
```



下面代码中 A B 两处应该怎么修改才能顺利编译？

```
func main() {
    var m map[string]int        //A
    m["a"] = 1
    if v := m["b"]; v != nil {  //B
        fmt.Println(v)
    }
}
```

```
答案：
A处：
m := make(map[sring]int)
B处：
if v, ok := m["b"] ; ok {
A处只是声明了m变量，而没有为该变量分配内存空间，不能直接赋值。

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
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowB())
    fmt.Println(b.ShowA())
}
```

- A. 23 13
- B. compilation error

```
答案： B

```

