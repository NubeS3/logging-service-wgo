package common

type ReqLog struct {
	Method   string `json:"method"`
	Req      string `json:"req"`
	SourceIp string `json:"source_ip"`
}

type UnauthReqLog struct {
	Event
	ReqLog
}

type AuthReqLog struct {
	Event
	ReqLog
	UserId string `json:"user_id"`
}

type AccessKeyReqLog struct {
	Event
	ReqLog
	Key string `json:"key"`
}

type SignedReqLog struct {
	Event
	ReqLog
	Public string `json:"public"`
}
