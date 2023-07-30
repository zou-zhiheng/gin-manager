package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sdlManager-mysql/global"
	"sdlManager-mysql/model"
)

const MysqlURI = "admin:admin@tcp(127.0.0.1:3306)/sdlManager"

func MysqlInit() {
	var err error
	if global.MysqlClientManager == nil {
		global.MysqlClientManager, err = gorm.Open("mysql", MysqlURI)
		if err != nil {
			fmt.Println("mysql连接失败", err)
		}

		//使用结构体默认名称作为表名
		global.MysqlClientManager.SingularTable(true)

		fmt.Println("mysql连接成功!")
	}

	//创建表
	{
		//用户
		if !global.MysqlClientManager.HasTable(&model.User{}) {
			if err = global.MysqlClientManager.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{}).Error; err != nil {
				panic(err)
			}
		}

		//API
		if !global.MysqlClientManager.HasTable(&model.Api{}) {
			if err = global.MysqlClientManager.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Api{}).Error; err != nil {
				panic(err)
			}
		}

		//角色
		if !global.MysqlClientManager.HasTable(&model.Role{}) {
			if err = global.MysqlClientManager.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Role{}).Error; err != nil {
				panic(err)
			}
		}
	}

	//后台管理
	{
		//用户
		global.UserTable = global.MysqlClientManager.Table("user")
		//API
		global.ApiTable = global.MysqlClientManager.Table("api")
		//角色
		global.RoleTable = global.MysqlClientManager.Table("role")
	}

}
