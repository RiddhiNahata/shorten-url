package utils

import (
	"errors"
	"sort"

	"github.com/RiddhiNahata/shorten-url/app/common/model"
	"gorm.io/gorm"
)

type kv struct {
	Key   string
	Value int
}

func CreateUser(db *gorm.DB, email string) (int64, error) {

	existingUser := &model.Users{}
	var userId int64 = 0
	db.First(&existingUser, "users.email = ?", email)

	if existingUser.UserId > 0 {
		userId = existingUser.UserId
	} else {
		var user model.Users
		user.Email = email
		// insert new db entry
		if result := db.Create(&user); result.Error != nil {
			return 0, errors.New(result.Error.Error())
		}
		userId = user.UserId
	}

	return userId, nil

}

func GetDomainCountByUserID(db *gorm.DB, userid int) ([]kv, error) {

	var urls []model.Urls

	domainMap := make(map[string]int)

	if userid == 0 {
		if result := db.Find(&urls); result.Error != nil {
			return nil, errors.New(result.Error.Error())
		}
	} else {
		if result := db.Where("urls.user_id = ?", userid).Find(&urls); result.Error != nil {
			return nil, errors.New(result.Error.Error())
		}
	}

	for _, url := range urls {

		if val, ok := domainMap[url.Domain]; ok {
			domainMap[url.Domain] = val + int(url.Count)
		} else {
			domainMap[url.Domain] = int(url.Count)
		}

	}
	// Sorting of map by values
	var keyVal []kv
	for k, v := range domainMap {
		keyVal = append(keyVal, kv{k, v})
	}

	sort.Slice(keyVal, func(i, j int) bool {
		return keyVal[i].Value > keyVal[j].Value
	})

	return keyVal, nil
}
