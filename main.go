package main

import (
	"fmt"
	"sdlManager-mysql/initialize"
	"sdlManager-mysql/router"
)

func init() {
	initialize.Init()
}

func main() {
	fmt.Println("s coding")
	engine := router.GetEngine()
	if err := engine.Run(":8060"); err != nil {
		panic(err)
	}

}
