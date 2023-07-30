package global

import (
	"github.com/jinzhu/gorm"
)

var (
	MysqlClientManager *gorm.DB
	UserTable          *gorm.DB
	ApiTable           *gorm.DB
	RoleTable          *gorm.DB
)
