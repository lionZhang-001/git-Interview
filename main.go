/*package main


func main(){

	DeferCalldd()
}
*/
/*package main

import "fmt"

func main() {
a := []int{7, 8, 9}
fmt.Printf("%+v\n", a)
ap(a)
fmt.Printf("%+v\n", a)
app(a)
fmt.Printf("%+v\n", a)

}

func ap(a []int) {
a = append(a, 10)
}

func app(a []int) {
a[0] = 1
}*/
package main

import (
	"fmt"
)

func hello() []string {
	return nil
}

func main() {
	h := hello
	fmt.Printf("%v\n" , h)
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}