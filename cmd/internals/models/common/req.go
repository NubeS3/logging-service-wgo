package common

type Req struct {
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Type   string   `json:"type"`
	Data   []string `json:"data"`
}
