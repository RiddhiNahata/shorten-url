package model

type Users struct {
	UserId int64  `json:"userid" gorm:"primaryKey"`
	Email  string `json:"email"`
}
