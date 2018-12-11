package utility

import (
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// ReadResponse takes a response value resulting from a http.Get and will
// use the Reader.Read function to read incomming bytes.  The incomming bytes
// will then be appended as a string and returned
// @params response - *http.Response
// @return string
func ReadResponse(response *http.Response) string {
	byteArray := make([]byte, 100)
	var dataString string
	for {
		n, err := response.Body.Read(byteArray)
		dataString += string(byteArray[:n])
		if err == io.EOF {
			break
		}
	}
	return dataString
}

func DecodeReadResponse(dataString, value string) string {
	var jsonParser map[string]interface{}

	bStream := []byte(dataString)

	err := json.Unmarshal(bStream, &jsonParser)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fContent, _ := jsonParser[value].(string)

	stringDecoded, _ := b64.StdEncoding.DecodeString(fContent)

	return ("\n" + string(stringDecoded))
}
