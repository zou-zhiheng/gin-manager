package main

import (
	"fmt"
	"sdlManager-mysql/router"
)

func init() {
	//initialize.Init()
}

func main() {
	fmt.Println("coding coding")
	engine := router.GetEngine()
	if err := engine.Run(":8060"); err != nil {
		panic(err)
	}

}
