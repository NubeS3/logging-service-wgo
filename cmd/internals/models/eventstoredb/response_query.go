package eventstoredb

import (
	"log"
	"time"

	goes "github.com/jetbasrawi/go.geteventstore"
)

var (
	eventActivators map[string]func() (interface{}, interface{})
)

func Query(streamName string, filter interface{}) {
	reader := client.NewStreamReader(streamName)
	for reader.Next() {
		if reader.Err() != nil {
			switch err := reader.Err().(type) {
			case *goes.ErrNotFound:
				<-time.After(time.Duration(10) * time.Second)
			case *goes.ErrNoMoreEvents:
				return
			default:
				log.Fatal(err)
			}
		} else {
			var event interface{}
			var meta interface{}
			if fn, ok := eventActivators[reader.EventResponse().Event.EventType]; ok {
				event, meta = fn()
				if err := reader.Scan(&event, &meta); err != nil {
					log.Fatal(err)
				}
				log.Printf(" - Event %d returned %#v\n Meta returned %#v\n", reader.EventResponse().Event.EventNumber, event, meta)
			} else {
				log.Fatalf("Error: Could not instantiate event of type %s", reader.EventResponse().Event.EventType)
			}
		}
	}
}
