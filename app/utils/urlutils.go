package utils

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/RiddhiNahata/shorten-url/app/common/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var Hostname string = "http://localhost:3000/"

func CreateNewShortUrl(db *gorm.DB, largeUrl string, userid int64) (model.Urls, error) {

	shortUrl := model.Urls{}
	urlList := []model.Urls{}

	db.Where("urls.large_url = ?", largeUrl).Find(&urlList)

	for _, url := range urlList {

		if url.UserId == userid {
			shortUrl = url
			break
		}

	}

	if shortUrl.Id == 0 {
		shortUrl.LargeUrl = largeUrl
		shortUrl.ShortUrl = Hostname + uuid.New().String()[:8]
		shortUrl.UserId = userid
		shortUrl.Domain = GetDomainName(largeUrl)
		shortUrl.Count = 0

		// insert new db entry
		if result := db.Create(&shortUrl); result.Error != nil {
			return shortUrl, errors.New(result.Error.Error())
		}
	}

	return shortUrl, nil
}

func GetLargeUrlFromShort(db *gorm.DB, shorturl string) (model.Urls, error) {
	existingUrl := model.Urls{}
	if result := db.First(&existingUrl, "urls.short_url = ?", shorturl); result.Error != nil {
		return existingUrl, errors.New(result.Error.Error())
	}
	return existingUrl, nil
}

func GetDomainName(largeUrl string) string {
	url, err := url.Parse(largeUrl)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(url.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	return domain
}

func UpdateURLCounter(db *gorm.DB, existingUrl model.Urls) error {

	existingUrl.Count = existingUrl.Count + 1
	if result := db.Save(&existingUrl); result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func GetUrlsByUserID(db *gorm.DB, userid int) ([]model.Urls, error) {

	var urls []model.Urls

	if userid == 0 {
		if result := db.Find(&urls); result.Error != nil {
			return urls, errors.New(result.Error.Error())
		}
	} else {
		if result := db.Where("urls.user_id = ?", userid).Find(&urls); result.Error != nil {
			return urls, errors.New(result.Error.Error())
		}
	}

	return urls, nil
}

func DeleteUrls(db *gorm.DB, id int) error {
	var url model.Urls
	if result := db.Where("id = ?", id).Delete(&url); result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}
