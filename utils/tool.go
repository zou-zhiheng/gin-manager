package utils

import (
	"encoding/json"
	"regexp"
	"strconv"
	"time"
)

func GetPage(currPage, pageSize string) (int64, int64, error) {
	var curr, size, skip int64
	curr, err := strconv.ParseInt(currPage, 10, 64)
	if err != nil {
		return skip, size, err
	}
	size, err = strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return skip, size, err
	}

	skip = (curr - 1) * size

	return skip, size, nil
}

func Regexp(str, s string) []string {
	//定义正则表达式
	regexpCompile := regexp.MustCompile(str)
	//使用正则表达式找与之相匹配的字符串，返回一个数组包含子表达式匹配的字符串
	return regexpCompile.FindStringSubmatch(s)
}
func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func Paging(data interface{}, pageSize, currPage int64) []interface{} {
	marshal, _ := json.Marshal(data)
	var res []interface{}
	_ = json.Unmarshal(marshal, &res)
	skip := int((currPage - 1) * pageSize)
	n := len(res)
	//手写分页查
	//不输入获取全部,或者inStocks为空
	if pageSize == 0 && currPage == 1 || data == nil {
		return res
	} else if skip > n { //起点大于数组长度
		return []interface{}{}
	} else if skip < n && int(pageSize*currPage) > n { //起点小于数组长度，但是总数大于
		return res[skip:n]
	} else {
		return res[skip : skip+int(pageSize)]
	}
}
func GetWeekDay(now time.Time) (string, string, time.Time, time.Time) {
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59", firstOfWeek, lastOfWeeK
}

// 获取当前月的时间函数封装
func GetMonthDay(now time.Time) (string, string) {
	//now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	//f := firstOfMonth.Unix()
	//l := lastOfMonth.Unix()
	//return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"

	//上面注释掉的特么多此一举，稍微优化一下这个函数
	return firstOfMonth.Format("2006-01-02") + " 00:00:00", lastOfMonth.Format("2006-01-02") + " 23:59:59"
}
