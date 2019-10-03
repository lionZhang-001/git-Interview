下面这段代码输出什么？

```
func main() {  
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
```

- A.13.1  
- B.13
- C.compilation error

```
答案：C
两个不同类型的数值不能相加，编译报错；go语言不允许隐式类型转换！！！！！
```



下面这段代码输出什么？

```
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{1, 2, 3, 4, 5}
    t := a[3:4:4]
    fmt.Println(t[0])
}
```

- A.3
- B.4
- C.compilation error  

```
答案： B
使用已有的切片/数组创建一个新的切片，方法有两种：
1: [i:j]表示从原数组的索引i开始，到索引j-1结束，作为新切片的内容。如果i省略，默认是从索引0开始，如果j省略，则默认是原切片或数组的长度。对于新切片的容量，当原数组或切片的容量是l时，则新切片的容量是l-i。
2: [i;j;k]第三个参数限定了新切片的容量。
```



下面这段代码输出什么？

```
func main() {
    a := [2]int{5, 6}
    b := [3]int{5, 6}
    if a == b {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}
```

- A. compilation error  
- B. equal  
- C. not equal  

```
答案： B
go中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的一部分，所以a和b是两种不同的类型。
```































