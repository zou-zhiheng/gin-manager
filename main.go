package main

import (
	"fmt"
	"sdlManager-mysql/router"
)

//
//func init() {
//	initialize.Init()
//}

func main() {

	fmt.Println("learning")
	engine := router.GetEngine()
	if err := engine.Run("7001"); err != nil {
		panic(err)
	}

}
