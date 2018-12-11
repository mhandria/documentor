package utility

import (
	"log"
	"net/http"
	"os"
	"io"
	b64 "encoding/base64"
	"encoding/json"
)

type FileFetcher struct {

}

func GetFileContent(url string) string {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	dataString := readResponse(response)
	
	var jsonParser map[string]interface{}

	err = json.Unmarshal([]byte(dataString), &jsonParser)
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	fContent, _ := jsonParser["content"].(string)

	stringDecoded, _ := b64.StdEncoding.DecodeString(fContent)

	return string(stringDecoded)
}

func readResponse(response *http.Response) string {
	byteArray := make([]byte, 100)
	var dataString string
	for {
		n, err := response.Body.Read(byteArray)
		if err == io.EOF {
			break
		}
		dataString += string(byteArray[:n])
	}
	return dataString
}
