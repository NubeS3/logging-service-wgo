package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

func UserLog(data interface{}) error {
	event := goes.NewEvent("user-"+goes.NewUUID(), "userEvent", data, nil)
	writer := client.NewStreamWriter("userStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

// func UserCreatedLog(data interface{}) error {
// 	event := goes.NewEvent("user-"+goes.NewUUID(), "userCreate", data, nil)
// 	writer := client.NewStreamWriter("userStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UserLogInLog(data interface{}) error {
// 	event := goes.NewEvent("user-"+goes.NewUUID(), "userLogIn", data, nil)
// 	writer := client.NewStreamWriter("userStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UserUpdateInfoLog(data interface{}) error {
// 	event := goes.NewEvent("user-"+goes.NewUUID(), "userUpdate", data, nil)
// 	writer := client.NewStreamWriter("userStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UserDeletedLog(data interface{}) error {
// 	event := goes.NewEvent("user-"+goes.NewUUID(), "userDeleted", data, nil)
// 	writer := client.NewStreamWriter("userStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UserUploadLog(data interface{}) error {
// 	event := goes.NewEvent("user-"+goes.NewUUID(), "userUpload", data, nil)
// 	writer := client.NewStreamWriter("userStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UserDownloadLog(data interface{}) error {
// 	event := goes.NewEvent("user-"+goes.NewUUID(), "userDownload", data, nil)
// 	writer := client.NewStreamWriter("userStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
