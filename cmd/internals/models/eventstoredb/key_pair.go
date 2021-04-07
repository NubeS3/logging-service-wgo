package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

func AccessKeyLog(data interface{}) error {
	event := goes.NewEvent("user-"+goes.NewUUID(), "accessKeyEvent", data, nil)
	writer := client.NewStreamWriter("keyStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

func KeyPairLog(data interface{}) error {
	event := goes.NewEvent("user-"+goes.NewUUID(), "keyPairEvent", data, nil)
	writer := client.NewStreamWriter("keyPairStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}
