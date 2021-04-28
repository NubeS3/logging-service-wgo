package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

func FileUploadedLog(data interface{}) error {
	event := goes.NewEvent("file-"+goes.NewUUID(), "fileUploaded", data, nil)
	writer := client.NewStreamWriter("fileStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

func FileDownloadedLog(data interface{}) error {
	event := goes.NewEvent("file-"+goes.NewUUID(), "fileDownloaded", data, nil)
	writer := client.NewStreamWriter("fileStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

func FileStagingLog(data interface{}) error {
	event := goes.NewEvent("file-"+goes.NewUUID(), "fileStaging", data, nil)
	writer := client.NewStreamWriter("fileStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

func FileUploadedSuccessLog(data interface{}) error {
	event := goes.NewEvent("file-"+goes.NewUUID(), "fileUploadedSuccess", data, nil)
	writer := client.NewStreamWriter("fileStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

// func FileDeletedLog(data interface{}) error {
// 	event := goes.NewEvent("file-"+goes.NewUUID(), "fileDeleted", data, nil)
// 	writer := client.NewStreamWriter("fileStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func FolderLog(data interface{}) error {
	event := goes.NewEvent("user-"+goes.NewUUID(), "folderEvent", data, nil)
	writer := client.NewStreamWriter("folderStream")
	err := writer.Append(nil, event)
	if err != nil {
		return err
	}

	return nil
}

// func FolderCreatedLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderCreated", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FolderDeletedLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderDelete", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FolderItemAddedLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderItemAdded", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FolderItemRemovedLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderItemRemoved", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FolderGotLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderGot", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FolderUploadedLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderUploaded", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func FolderDownloadedLog(data interface{}) error {
// 	event := goes.NewEvent("folder-"+goes.NewUUID(), "folderDownloaded", data, nil)
// 	writer := client.NewStreamWriter("folderStream")
// 	err := writer.Append(nil, event)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
