package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	Id        string `gorm:"PrimaryKey"`
	Writer    string `valid:"required~your writer is required"`
	Title     string `valid:"required~your title is required"`
	Content   string `valid:"required~your content is required"`
	Date      string `valid:"required~your date is required, date~invalid date format"`
	Image     string `valid:"required~your image is required, image~invalid image format"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ArticleReq struct {
	Writer  string `json:"writer" form:"writer"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Date    string `json:"date" form:"date"`
	Image   string `json:"image" form:"image"`
}

type ArticleResp struct {
	Id      string `json:"id"`
	Writer  string `json:"writer"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Image   string `json:"image"`
}

func (a *Articles) BeforeCreate(tx *gorm.DB) error {
	a.Id = helpers.CreateId()
	return nil
}
