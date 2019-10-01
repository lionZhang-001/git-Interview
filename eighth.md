关于init函数，下面说法正确的是()

- A. 一个包中，可以包含多个 init 函数；
- B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
- C. main 包中，不能有 init 函数；
- D. init 函数可以被其他函数调用

```
答案： A B 
关于init函数有几个要注意的地方：
1. init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;

2. 一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；

3. 同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的（看下图）;

4. init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;

5. 一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；

6. 引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；
```

![640](/home/zyy/typora-notebook/编程语言/go/go培训/picture/640.webp)





下面这段代码输出什么以及原因？

```
func hello() []string {  
     return nil
}
 
func main() {  
    h := hello
    if h == nil {
        fmt.Println("nil")
        } else {
        fmt.Println("not nil")
    }
}
```

- A. nil
- B. not nil
- C. compilation error  

```
答案： B
这是把函数赋值给h，而不是把返回值赋值给h，我记得函数是引用类型，是函数的地址。
```



下面这段代码能否编译通过？如果可以，输出什么？

```go
 func GetValue() int {
     return 1
 }
 
 func main() {
     i := GetValue()
     switch i.(type) {
     case int:
         println("int")
    case string:
        println("string")
        case interface{}:
        println("interface")
    default:
        println("unknown")
    }
}
```

```
答案：不能编译通过
类型断言：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型断言。
```

