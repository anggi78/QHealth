package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	Id        string         `gorm:"PrimaryKey" json:"id"`
	Writer    string         `json:"writer" valid:"required~your writer is required"`
	Title     string         `json:"title" valid:"required~your title is required"`
	Content   string         `json:"content" valid:"required~your content is required"`
	Date      string         `json:"date" valid:"required~your date is required, date~invalid date format"`
	Image     string         `json:"image" valid:"required~your image is required, image~invalid image format"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
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
