package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once //once.Do(func(){})其中func只执行一次
var lock = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func newSingle() {
	singleInstance = &Single{}
}

func GetInstance() *Single {
	if singleInstance == nil {
		//once.Do(
		//	func() {
		//		fmt.Println("instance is creating")
		//		newSingle() //只会执行一次
		//	})
		//不放在外层循环是因为，只有在创建时需要同步，而创建之后它！=nil，不用浪费事件加锁
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("instance is creating")
			newSingle()
		} else {
			fmt.Println("instance has been created")
		}
	} else {
		fmt.Println("instance has been created")
	}
	return singleInstance
}

func (s *Single) DoSomething() {
	fmt.Printf("%p\n", s)
}
