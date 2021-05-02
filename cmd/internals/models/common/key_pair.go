package common

type KeyPairLog struct {
	Event
	BucketId     string `json:"bucket_id"`
	Public       string `json:"public"`
	GeneratorUid string `json:"uid"`
}
