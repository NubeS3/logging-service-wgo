package common

type ErrLog struct {
	Event
	Error string `json:"content"`
}
