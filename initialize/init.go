package initialize

func Init() {
	MysqlInit()
	RedisInit()
	RabbitMQInit()
	SnowFlakeInit()
}
