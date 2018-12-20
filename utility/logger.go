package utility

import (
	"encoding/json"
	"log"
)

// LogIfError logs the error into a json format and prints it back
// as with log.Fatal.  This function will return true if error != nil
// and return false if error is nil
func LogIfError(err error, extraInfo ...interface{}) (isError bool) {
	isError = (err != nil)
	if isError {
		event, _ := json.Marshal(logEvent{Message: err.Error(), Extra: extraInfo})
		log.Fatal(string(event))
	}
	return
}

type logEvent struct {
	Message string
	Extra   []interface{}
}
