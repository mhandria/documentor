package main 

import (
	"fmt"
	"net/http"
	"os"
	"io"
b64 "encoding/base64"
	"encoding/json"
)

func main() {

	resp, err := http.Get("https://api.github.com/repos/michael-handria/vue-test-project/git/blobs/ade22e26b02d5a056f820437293d9e9e13d9aa66")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// dat, err := ioutil.ReadFile("https://api.github.com/repos/michael-handria/vue-test-project/git/blobs/ade22e26b02d5a056f820437293d9e9e13d9aa66")
	// if err != nil {
	// 	fmt.Printf("well that didn't work :( %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Print(string(dat))
	b := make([]byte, 1000)
	n, err := resp.Body.Read(b)
	if err != io.EOF {
		fmt.Println("read data to small ")
		os.Exit(1) 
	}
	var bd map[string]interface{}

	err = json.Unmarshal(b[:n], &bd)

	if err != nil {
		fmt.Printf("oops something happened when parsing json\n %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("body: %q\n", b[:n])
	str, _ := bd["content"].(string)

	sDec, _ := b64.StdEncoding.DecodeString(str)
	//
	fmt.Println(string(sDec))

	fmt.Println("Hello Blizzard!")
}