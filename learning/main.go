package main

/*doc*/
import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
)

func test0() {
	var a string
	var b int
	var c bool

	a = "6"

	d := 6
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func test() {
	x, y, z := 'c', 1.9, "str"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func test2() {
	var (
		a int
		b float64
		c string
	)
	const (
		x bool = true
		y bool = false
		z bool = true
	)
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
}

func test3() {
	fmt.Print("hello", "world", "\n")
	fmt.Println("hello", "world")
	fmt.Printf("type %T, data %#v\n", true, true)
	fmt.Printf("type %T data %#v\n", '1', '1')
	fmt.Printf("%%b %b\n", 8)
	fmt.Printf("%%+d %+d\n", 8)
	fmt.Printf("%q\n", "string")
	fmt.Printf("%x\n", "string")
	fmt.Printf("% x\n", "string")
}

func test5() {
	a := 1
	b := -1
	var c uint8 = 128
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %d\n", b, b)
	fmt.Printf("%T %d\n", c, c)
}

func test6() {
	var x = [...]int{1, 2, 3, 4}
	fmt.Println("len(x)=", len(x))

	y := [...]string{0: "I'm", 9: "Iron", 20: "man"}
	fmt.Println(y)
}

//slice
func test7() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array := [...]int{1, 2, 3, 4, 5}
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	fmt.Printf("%T, %v\n", slice1, slice1)
	fmt.Printf("%T, %v\n", array, array)
	fmt.Printf("slice3: %v\n", slice3)
}

func test8() {
	subject := []string{"math", "cs", "chinese", "english"}
	for _, name := range subject {
		fmt.Println(name)
	}
}

func test9(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("string type : %v", v)
	case int:
		fmt.Printf("this is int type: %v", v)
	default:
		fmt.Println(v)
	}
}

type Person struct {
	name  string
	age   int
	agend string
}

func (p Person) String() string {
	return fmt.Sprintf("name = %s, age = %d, agend = %s", p.name, p.age, p.agend)
}

func test10() {
	p := Person{"小黑", 19, "male"}

	fmt.Println(p)
}

func test11() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func test12() {
	ch := make(chan string, 2)
	ch <- strconv.Itoa(1123)
	ch <- strconv.Itoa(54231)
	fmt.Printf("%q\n", <-ch)
	fmt.Printf("%q\n", <-ch)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
