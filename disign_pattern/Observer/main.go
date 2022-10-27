package main

import (
	"Observer/Observer"
	"fmt"
)

func main() {
	publisher := &Observer.Publisher{}
	publisher.AddListener(&EmailListener{})
	publisher.AddListener(&LogListener{})
	publisher.Notify()
}

type EmailListener struct {
}

func (e *EmailListener) Update() {
	fmt.Println("email 接收到通知")
}

type LogListener struct {
}

func (l *LogListener) Update() {
	fmt.Println("logg 接收到通知")
}
