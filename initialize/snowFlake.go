package initialize

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sdlManager-mysql/global"
)

// SnowFlakeInit 初始化雪花算法
func SnowFlakeInit() {
	if global.SnowFlake == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			fmt.Println("snowflake init:", err)
		}

		global.SnowFlake = node
	}
}
