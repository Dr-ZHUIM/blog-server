package DTOs

type Article struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Artist    string `json:"artist"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
}
