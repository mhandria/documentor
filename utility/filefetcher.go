//Package provides function to get file content given url
package utility

import (
	"log"
	"net/http"
	"os"
)

func GetGithubFileContent(url string) string {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dataString := ReadResponse(response)

	return DecodeReadResponse(dataString, "content")
}
