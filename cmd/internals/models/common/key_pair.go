package common

type KeyPairLog struct {
	Event
	BucketId string `json:"bid"`
	UserId string `json:"uid"`
	
}
