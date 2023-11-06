package DTOs

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int64  `json:"id"`
}
