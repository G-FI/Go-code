package bridge

import "fmt"

type timerPrinter struct {
	time string
}

func newTimerPrinter(time string) *timerPrinter {
	return &timerPrinter{time: time}
}
func (t *timerPrinter) PrintFile() {
	fmt.Printf("timestap：%s\n", t.time)
	fmt.Println("timerPrinter is printing file")
}
