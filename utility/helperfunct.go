//Package provides function that helps with documentor
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

// DecodeReadResponse takes a json string and a value within that json string
// that the user wants to decode.
// NOTE: this will only decode a base64 encryption
// @param dataString - json string
// @param value - content string that needs to be decoded
// @return decoded string
func DecodeReadResponse(dataString, value string) string {
	var jsonParser map[string]interface{}

	bStream := []byte(dataString)

	err := json.Unmarshal(bStream, &jsonParser)

	if err != nil {
		log.Fatal(err)
		os.Exit(1) //TODO: remove later change to goroutine
	}

	fContent, _ := jsonParser[value].(string)

	stringDecoded, _ := b64.StdEncoding.DecodeString(fContent)

	return ("\n" + string(stringDecoded))
}

func StringToJsonArray(dataString string) [](map[string]interface{}) {
	var jsonParsed []map[string]interface{}
	bStream := []byte(dataString)
	err := json.Unmarshal(bStream, &jsonParsed)
	if err != nil {
		log.Fatal(err)
		os.Exit(1) //TODO: remove later change to goroutine
	}

	return jsonParsed
}

func TryRequest(req *http.Request, err error) *http.Request {
	if err != nil {
		log.Fatal(err)
		os.Exit(1) //TODO: remove later change to goroutine
	}
	return req
}

func TryResponse(res *http.Response, err error) *http.Response {
	if err != nil {
		log.Fatal(err)
		os.Exit(1) //TODO: remove later change to goroutine
	}
	return res
}
