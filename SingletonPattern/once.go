// 使用 sync.Once 来创建对象

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	onceInstance *single
	once         sync.Once
)

func getOnceInstance() *single {
	if onceInstance == nil {
		once.Do(func() {
			fmt.Println("creating multi-thread safe once instance now.")
			onceInstance = &single{}
		})
	} else {
		fmt.Println("multi-thread safe once instance already created.")
	}
	return onceInstance
}

func multiThreadSafeOnce() {
	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				getOnceInstance()
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
