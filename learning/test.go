package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

func foo() *[]int {
	v := make([]int, 10, 100)
	v[1] = 8
	return &v
}

func negative(v *[]int) {
	for i := 0; i < len(*v); i++ {
		(*v)[i] = -(*v)[i]
	}
}

type varI interface {
	out()
}
type myout struct {
	msg string
}
type Integer int

func (self Integer) out() {
	fmt.Println(self)
}

func (self myout) out() {
	fmt.Println(self.msg)
}

// func main() {
// 	ch := make(chan string)

// 	go sendData(ch)
// 	go getData(ch)

// 	//time.Sleep(1e9)
// }

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
