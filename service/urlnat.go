package service

import (
	"time"
	"urlnat/global"
	"urlnat/model"
)

func AddUrlNat(data model.UrlNat) int64 {
	return global.DB.Create(&data).RowsAffected
}

func DeleteUrlNat(data model.UrlNat) int64 {
	return global.DB.Where("`key` = ?", data.Key).Delete(&data).RowsAffected
}

func UpdateUrlNat(data model.UrlNat) int64 {
	return global.DB.Where("`key` = ?", data.Key).Updates(&data).RowsAffected
}

func GetUrlNat(code string) model.UrlNat {
	var data model.UrlNat
	global.DB.Where("`code` = ?", code).First(&data)
	return data
}

func ClearUrlNat() int64 {
	now := time.Now()
	date := now.Format("2006-01-02")
	return global.DB.Where("`date` < ?", date).Delete(&model.UrlNat{}).RowsAffected
}
