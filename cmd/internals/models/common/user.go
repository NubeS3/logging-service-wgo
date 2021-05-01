package common

type UserLog struct {
	Event
	UserId    string `json:"content"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Pass      string `json:"password"`
	IsActive  bool   `json:"is_active"`
	IsBanned  bool   `json:"is_banned"`
}
