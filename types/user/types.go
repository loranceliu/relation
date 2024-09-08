package user

type UserRequest struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Status   int8   `json:"status"`
}
