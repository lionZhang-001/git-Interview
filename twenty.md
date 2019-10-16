下面这段代码正确的输出是什么？

```
func f() {
    defer fmt.Println("D")
    fmt.Println("F")
}

func main() {
    f()
    fmt.Println("M")
}
```

- A. F M D
- B. D F M
- C. F D M

```
答案： C

```



下面代码输出什么？

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

    person = &Person{29}
}
```

```
答案：
29
28
28
解析：
1: 28
2:28
3:29
```







