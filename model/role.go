package model

type Role struct {
	Id         int64   `db:"id" json:"id" gorm:"column:id;type:bigint;primary_key;not null"`
	Name       string  `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Code       string  `db:"code" json:"code" gorm:"column:code;type:varchar(10);not null"`
	Api        string  `db:"api" json:"api" gorm:"column:api;type:text"`
	Apis       []int64 `db:"apis" json:"apis" gorm:"-"`
	RoleRoutes string  `db:"roleRoutes" json:"roleRoutes" gorm:"column:roleRoutes;type:varchar(20)"`
	FirstPage  string  `db:"firstPage" json:"firstPage" gorm:"column:firstPage;type:varchar(20)"`
	Desc       string  `db:"desc" json:"desc" gorm:"column:desc;type:varchar(255)"`
	CreateTime string  `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
	UpdateTime string  `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:varchar(20)"`
}
