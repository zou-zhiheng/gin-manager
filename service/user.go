package service

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"sdlManager-mysql/global"
	"sdlManager-mysql/middleware"
	"sdlManager-mysql/model"
	"sdlManager-mysql/utils"
	"strconv"
	"time"
)

func CreateUser(user model.User) utils.Response {

	if user.Account == "" || user.Name == "" || user.Password == "" || user.Phone == "" {
		return utils.ErrorMess("失败，重要参数缺失", nil)
	}

	var userDB []model.User
	res := global.UserTable.Where("isValid = true and account = ?", user.Account).Find(&userDB)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error)
	}

	if res.RowsAffected != 0 {
		return utils.ErrorMess("该用户已存在", nil)
	}

	//根据时间戳生成种子，防止恶意伪造
	rand.Seed(time.Now().Unix())

	//生成盐
	user.Salt = strconv.FormatInt(time.Now().Unix(), 10)
	//密码加密加盐
	encryptedPass, er := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
	if er != nil {
		return utils.ErrorMess("密码加密失败", er.Error())
	}

	user.Password = string(encryptedPass)
	user.CreateTime = utils.TimeFormat(time.Now())
	user.Id = global.SnowFlake.Generate().Int64()
	if len(user.RoleIds) != 0 {
		temp, err := json.Marshal(user.RoleIds)
		if err != nil {
			return utils.ErrorMess("失败", nil)
		}

		user.RoleId = string(temp)
	}

	res = global.UserTable.Create(&user)
	if res.Error != nil {
		return utils.ErrorMess("创建失败", res.Error.Error())
	}

	return utils.SuccessMess("成功", res.RowsAffected)

}

func UpdateUser(user model.User) utils.Response {

	if user.Name == "" {
		return utils.ErrorMess("失败，重要参数缺失", nil)
	}

	var userDB model.User
	res := global.UserTable.Where("isValid = true and  id = ?", user.Id).Find(&userDB)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error)
	}

	user.Account = userDB.Account
	user.Password = userDB.Password
	user.Phone = userDB.Phone
	user.Salt = userDB.Salt
	user.UpdateTime = utils.TimeFormat(time.Now())
	temp, err := json.Marshal(user.RoleIds)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	user.RoleId = string(temp)

	res = global.UserTable.Where("id =?", user.Id).Updates(&user)

	return utils.SuccessMess("成功", res.RowsAffected)
}

func GetUser(name, currPage, pageSize, startTime, endTime string) utils.Response {

	skip, size, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err)
	}

	db := global.UserTable
	if startTime != "" && endTime != "" {
		db = global.UserTable.Where("createTime >= ? and createTime <= ?", startTime, endTime)
	}

	var count int
	var userDB []model.User
	res := db.Order("id desc").Where("isValid = true and name like ?", "%"+name+"%").Limit(size).Offset(skip).Find(&userDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}

	//json反序列化
	for i := range userDB {
		_ = json.Unmarshal([]byte(userDB[i].RoleId), &userDB[i].RoleIds)
	}

	return utils.SuccessMess("成功", struct {
		Count int          `json:"count" bson:"count"`
		Data  []model.User `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  userDB,
	})
}

func DeleteUser(idStr string) utils.Response {

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	var userDB model.User
	res := global.UserTable.Where("isValid = true and id = ?", id).Find(&userDB)
	if res.Error != nil {
		return utils.ErrorMess("失败，该用户不存在", res.Error)
	}

	res = global.UserTable.Where("id = ?", id).Update("isValid", false)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error)
	}

	return utils.SuccessMess("成功", res.RowsAffected)
}

func Login(user model.User) utils.Response {

	if user.Account == "" || user.Password == "" {
		return utils.ErrorMess("失败", nil)
	}
	var userDB model.User
	res := global.UserTable.Where("account = ?", user.Account).Find(&userDB)
	if res.Error != nil {
		return utils.ErrorMess("失败，该用户不存在", nil)
	}

	//密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password+userDB.Salt)); err != nil {
		return utils.ErrorMess("失败,密码错误", nil)
	}

	//json反序列化
	err := json.Unmarshal([]byte(userDB.RoleId), &userDB.RoleIds)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}

	//生成token
	token, err := middleware.CreateToken(userDB)
	if err != nil {
		return utils.ErrorMess("生成toke失败", err.Error())
	}

	data := map[string]interface{}{
		"id":      userDB.Id,
		"name":    userDB.Name,
		"account": userDB.Account,
		"roleId":  userDB.RoleId,
		"roleIds": userDB.RoleIds,
		"sex":     userDB.Sex,
		"token":   token,
	}

	return utils.SuccessMess("成功", data)
}
