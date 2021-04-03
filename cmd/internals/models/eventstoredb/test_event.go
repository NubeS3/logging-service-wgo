package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
	"log"
	"math/rand"
)

type testEvent struct {
	Val int
}

type testMetadata struct {
	W int
}

func SaveTestEvent() {
	testEvent := &testEvent{Val: rand.Int() % 10}
	testMetadata := &testMetadata{W: rand.Int() % 3}

	event := goes.NewEvent(goes.NewUUID(), "t", testEvent, testMetadata)
	writer := client.NewStreamWriter("tStream")
	err := writer.Append(nil, event)
	if err != nil {
		log.Fatal(err.Error()) // Handle errors
	}
}
