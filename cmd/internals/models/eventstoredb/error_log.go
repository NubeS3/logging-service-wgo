package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

func SaveErrorLog(data interface{}) error {
	event := goes.NewEvent(goes.NewUUID(), "error", data, nil)
	writer := client.NewStreamWriter("errorStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}
