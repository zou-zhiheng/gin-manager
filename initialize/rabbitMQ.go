package initialize

import (
	"sdlManager-mysql/global"
	"sdlManager-mysql/model"
)

// RabbitMQInit 初始化rabbitmq
func RabbitMQInit() {
	if global.RabbitmqSimple == nil {
		global.RabbitmqSimple = model.NewRabbitMQSimple("sttch")
	}
}
