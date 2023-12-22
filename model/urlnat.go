package model

type UrlNat struct {
	Code string `json:"code" gorm:"primary_key" binding:"required"`
	Note string `json:"note"`
	Url  string `json:"url"`
	Key  string `json:"key"`
	Date string `json:"date"`
}
