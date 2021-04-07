package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

func BucketLog(data interface{}) error {
	event := goes.NewEvent("user-"+goes.NewUUID(), "bucketEvent", data, nil)
	writer := client.NewStreamWriter("bucketStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}
