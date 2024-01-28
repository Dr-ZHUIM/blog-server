package DTOs

type Component struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Function string    `json:"function"`
	Articles []Article `json:"-" gorm:"many2many:article_components;"`
}

type ComponentRes struct {
	ID       int    `json:"id"`
	Function string `json:"function"`
}
