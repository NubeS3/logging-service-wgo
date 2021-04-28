package logic

import (
	"encoding/json"
	"fmt"
	"log-service-go/cmd/internals/models/common"
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
