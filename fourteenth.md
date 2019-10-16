下面的程序会输出什么？

```
func main() {
    str := "hello"
    str[0] = 'x'
    fmt.Println(str)
}
```

- A. hello
- B. xello
- C. compilation error

```
答案： C
我当时选的是B。这是一个基本的问题，在go语言中，字符串是常量，是只读的，无法通过下标的形式去修改它。如果向修改字符串，就要先强制转换成字节切片[]byte，在去修改。
```



下面代码输出什么？

```
func incr(p *int) int {
    *p++
    return *p
}

func main() {
    p :=1
    incr(&p)
    fmt.Println(p)
}
```

- A. 1
- B. 2
- C. 3

```
答案： B

```



对 add() 函数调用正确的是（）

```
func add(args ...int) int {

    sum := 0
    for _, arg := range args {
        sum += arg
    }
    return sum
}
```

- A. add(1, 2)
- B. add(1, 3, 7)
- C. add([]int{1, 2})
- D. add([]int{1, 3, 7}…)

```
答案：A B D
可变函数
```

