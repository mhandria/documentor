package utility

import (
	"encoding/json"
	"log"
)

func LogIfError(err error, extraInfo ...interface{}) bool {
	if err != nil {
		event, _ := json.Marshal(logEvent{Message: err.Error(), Extra: extraInfo})
		log.Fatal(string(event))
		return true
	}
	return false
}

type logEvent struct {
	Message string
	Extra   []interface{}
}
