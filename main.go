package main

import (
	"sdlManager-mysql/initialize"
	"sdlManager-mysql/router"
)

func init() {
	initialize.Init()
}

func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8060"); err != nil {
		panic(err)
	}

}
