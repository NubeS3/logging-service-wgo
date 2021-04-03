package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

func SaveErrorLog(data interface{}) error {
	event := goes.NewEvent(goes.NewUUID(), "error", data, nil)
	writer := client.NewStreamWriter("tStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}
