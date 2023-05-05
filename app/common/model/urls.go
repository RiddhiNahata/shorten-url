package model

type Urls struct {
	Id       int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	LargeUrl string `json:"largeurl"`
	ShortUrl string `json:"shorturl"`
	UserId   int64  `json:"userid"`
	Count    int64  `json:"count"`
	Domain   string `json:"domain"`
}
