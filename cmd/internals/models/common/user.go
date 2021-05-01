package common

type UserLog struct {
	Event
	UserId   string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
