package main

import (
	"fmt"
	"sdlManager-mysql/initialize"
	"sync"
	"time"
)

func init() {
	initialize.Init()
}

func main() {

	var pool sync.Pool
	pool.Put(func() {
		fmt.Println("pool")
	})

	a := pool.Get().(func())
	go a()

	time.Sleep(1)

}
