package DTOs

import (
	"fmt"

	"gorm.io/gorm"
)

type Article struct {
	ID           int         `json:"id" gorm:"primaryKey"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	Artist       string      `json:"artist"`
	Abstract     string      `json:"abstract"`
	Category     string      `json:"category"`
	Layer        string      `json:"layer"`
	SideCategory string      `json:"sidecategory"`
	CreatedAt    int64       `json:"createdat"`
	PathName     string      `json:"pathname"`
	Components   []Component `json:"components" gorm:"many2many:article_components;"`
}

type ArticleReview struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Abstract     string `json:"abstract"`
	Category     string `json:"category"`
	Layer        string `json:"layer"`
	SideCategory string `json:"sidecategory"`
	CreatedAt    int64  `json:"createdat"`
	PathName     string `json:"pathname"`
}

type KeyWord struct {
	Keyword string
}

func (article *Article) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("article artist: ", article.Artist)
	return nil
}
