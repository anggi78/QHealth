package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type ArticleView struct {
	Id        string `gorm:"PrimaryKey"`
	IdUser    string
	IdArticle string
	User      User      `gorm:"foreignKey:IdUser;references:Id"`
	Article   Articles  `gorm:"foreignKey:IdArticle;references:Id"`
	ViewedAt  time.Time 
}

type ArticleViewResp struct {
	Id        string        `json:"id"`
	IdUser    string        `json:"id_user"`
	User      UserResp      `json:"user"`
	Article   []ArticleResp `json:"article"`
	ViewedAt  time.Time     `json:"viewed_at"`
}

func (v *ArticleView) BeforeCreate(tx *gorm.DB) error {
	v.Id = helpers.CreateId()
	return nil
}
