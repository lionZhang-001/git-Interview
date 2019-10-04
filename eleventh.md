关于 cap() 函数的适用类型，下面说法正确的是()

- A. array
- B. slice
- C. map
- D. channel

```
答案： A B  D
根据go1.13官方文档显示，cap函数适用于以下几个类型：
array; pointer to array; slice ; channel
array：表示数组元素的个数，其值与len()函数相等
pointer to array:表示该指针所指数组元素的个数，其值与len()函数相等
slice：表示切片所能容纳的最大元素的个数，如果该切片为nil，则cap返回0
channel：表示当前chan的总容量，如果该channel为nil，则cap返回0

并不适用与map
```



下面这段代码输出什么？

```
func main() {  
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
```

- A. nil
- B. not nil
- C. compilation error  

```
答案： A
当且仅当接口的动态值和动态类型都为nil时，其接口类型才为nil。
```



下面这段代码输出什么？

```
func main() {  
    s := make(map[string]int)
    delete(s, "h")
    fmt.Println(s["h"])
}
```

- A. runtime panic
- B. 0
- C. compilation error 

```
答案：B

删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。
```











