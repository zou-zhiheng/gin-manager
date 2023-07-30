package service

import (
	"sdlManager-mysql/global"
	"sdlManager-mysql/model"
	"sdlManager-mysql/utils"
	"strconv"
	"time"
)

func GetApi(name, currPage, pageSize, startTime, endTime string) utils.Response {

	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	db := global.ApiTable
	if startTime != "" && endTime != "" {
		db = db.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}

	var count int64
	var apiDB []model.Api
	res := db.Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&apiDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", struct {
		Count int64       `json:"count" bson:"count"`
		Data  []model.Api `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  apiDB,
	})
}

func CreateApi(api model.Api) utils.Response {

	if api.Name == "" || api.Url == "" || (api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("失败,参数错误", nil)
	}

	var apiDB model.Api
	res := global.ApiTable.Where("name = ?", api.Name).Or("url = ? and method = ?", api.Url, api.Method).Find(&apiDB)
	if res.Error == nil {
		return utils.ErrorMess("失败，该API已存在", nil)
	}

	api.CreateTime = utils.TimeFormat(time.Now())
	api.Id = global.SnowFlake.Generate().Int64()

	//数据插入
	res = global.ApiTable.Create(&api)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", res.RowsAffected)

}

func UpdateApi(api model.Api) utils.Response {

	if api.Id == 0 || api.Name == "" || api.Url == "" || (api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("失败,参数错误", nil)
	}

	var apiDB model.Api
	res := global.ApiTable.Where("id = ?", api.Id).Find(&apiDB)
	if res.Error != nil {
		return utils.ErrorMess("失败，该API不存在", nil)
	}

	api.CreateTime = apiDB.CreateTime
	api.UpdateTime = utils.TimeFormat(time.Now())

	res = global.ApiTable.Where("id = ?", api.Id).Update(&api)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", res.RowsAffected)
}

func DeleteApi(idStr string) utils.Response {

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	res := global.ApiTable.Delete("id = ?", id)

	return utils.SuccessMess("成功", res.RowsAffected)
}
