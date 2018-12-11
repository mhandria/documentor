package main

import (
	utility "documentor/utility"
	"log"
	"net/http"
	"os"
)

func main() {
	val := utility.GetGithubFileContent("https://api.github.com/repos/michael-handria/vue-test-project/git/blobs/ade22e26b02d5a056f820437293d9e9e13d9aa66")
	log.Println(val)

	resp, err := http.Get("https://ghosthub.corp.blizzard.net/PlayerBehavior")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dataString := utility.ReadResponse(resp)
	log.Println(dataString)
}
