package service

import (
	"encoding/json"
	"sdlManager-mysql/global"
	"sdlManager-mysql/model"
	"sdlManager-mysql/utils"
	"strconv"
	"time"
)

func GetRole(name, currPage, pageSize, startTime, endTime string) utils.Response {

	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	db := global.RoleTable
	if startTime != "" && endTime != "" {
		db = db.Where("createTime >= ? and createTime <= ?")
	}

	var count int64
	var roleDB []model.Role
	res := db.Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&roleDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	//json反序列化
	for i := range roleDB {
		_ = json.Unmarshal([]byte(roleDB[i].Api), &roleDB[i].Apis)
	}

	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.Role `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  roleDB,
	})
}

func CreateRole(role model.Role) utils.Response {

	if role.Name == "" || role.Code == "" {
		return utils.ErrorMess("失败,重要参数缺失", nil)
	}
	var roleDB model.Role
	res := global.RoleTable.Where("name = ?", role.Name).Find(&roleDB)
	if res.Error == nil {
		return utils.ErrorMess("失败,该角色已存在", res.Error.Error())
	}

	role.Id = global.SnowFlake.Generate().Int64()
	role.CreateTime = utils.TimeFormat(time.Now())

	res = global.RoleTable.Create(&role)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", res.RowsAffected)

}

func UpdateRole(role model.Role) utils.Response {

	if role.Id == 0 || role.Code == "" {
		return utils.ErrorMess("失败", nil)
	}

	var roleDB []model.Role
	res := global.RoleTable.Where("id = ? or code = ?", role.Id, role.Code).Find(&roleDB)
	if res.Error != nil || len(roleDB) > 1 {
		return utils.ErrorMess("失败，该角色不存在", res.Error.Error())
	}

	role.CreateTime = roleDB[0].CreateTime
	role.UpdateTime = utils.TimeFormat(time.Now())

	temp, err := json.Marshal(role.Apis)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	role.Api = string(temp)

	res = global.RoleTable.Where("id = ?", role.Id).Update(&role)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", res.RowsAffected)
}

func DeleteRole(idStr string) utils.Response {

	if idStr == "" {
		return utils.ErrorMess("失败", nil)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	//角色查找
	var roleDB model.Role
	res := global.RoleTable.Where("id = ?", id).Find(&roleDB)
	if res.Error != nil {
		return utils.ErrorMess("该角色不存在", res.Error.Error())
	}

	//用户权限更新
	var userDB []model.User
	_ = global.UserTable.Where("roleId like ?", "%"+idStr+"%").Find(&userDB)
	if len(userDB) != 0 {
		for i := range userDB {
			//删除
			userDB[i].RoleId = utils.JsonDeleteOne(userDB[i].RoleId, idStr)
		}

		res = global.UserTable.Where("roleId like ?", "%"+idStr+"%").Updates(&userDB)
	}

	//角色删除
	res = global.RoleTable.Delete("id = ?", id)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", res.RowsAffected)
}
