package logic

import (
	"encoding/json"
	"fmt"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"time"
)

func TestErr() {
	testData := common.Req{
		Limit:  10,
		Offset: 0,
		Type:   "Date",
		Data:   []string{time.Now().Add(time.Hour * -1).String(), time.Now().Add(time.Hour * 1).String()},
	}

	jsonTestData, _ := json.Marshal(testData)

	res, err := nc.Request(errSubj+"query", jsonTestData, time.Second*10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res.Data))
}

func TestSendFile() {
	log := common.FileLog{
		Event: common.Event{
			Type: "Upload",
			Date: time.Now(),
		},
		FId:        "1234",
		FileName:   "test.file",
		Size:       1024,
		BucketId:   "12345",
		UploadDate: time.Now(),
		Uid:        "0",
	}

	jsonTestData, _ := json.Marshal(log)

	err := sc.Publish(fileSubj, jsonTestData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestFile() {
	testData := common.Req{
		Limit:  10,
		Offset: 0,
		Type:   "Date",
		Data:   []string{time.Now().Add(time.Hour * -1).Format(time.RFC3339), time.Now().Add(time.Hour * 1).Format(time.RFC3339)},
	}

	jsonTestData, _ := json.Marshal(testData)

	res, err := nc.Request(fileSubj+"query", jsonTestData, time.Second*10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res.Data))
}
