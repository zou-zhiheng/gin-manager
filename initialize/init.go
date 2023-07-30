package initialize

func Init() {
	MysqlInit()
	RedisInit()
	SnowFlakeInit()
}
