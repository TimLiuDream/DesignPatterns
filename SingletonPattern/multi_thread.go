package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

var multiThreadSafeInstance *single

func getMultiThreadSafeInstance() *single {
	if multiThreadSafeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if multiThreadSafeInstance == nil {
			fmt.Println("creating multi-thread safe instance now.")
			multiThreadSafeInstance = new(single)
		} else {
			fmt.Println("multi-thread safe instance already created.")
		}
	} else {
		fmt.Println("multi-thread safe instance already created.")
	}

	return multiThreadSafeInstance
}

func multiThreadSafe() {
	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				getMultiThreadSafeInstance()
			}
		}()
	}
	select {}
}
