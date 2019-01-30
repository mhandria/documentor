package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/mhandria/documentor/secret"
	"github.com/mhandria/documentor/utility"
)

const ghosthubKey, ghosthubAPIURL = "blank", "blank"

var client = &http.Client{}

//GetOrganizationMembers -> get organization members from github given a org name
func GetOrganizationMembers(orgName string) {
	endpoint := ghosthubAPIURL + "/orgs/" + orgName + "/members"
	req := utility.TryRequest(http.NewRequest("GET", endpoint, nil))

	req.Header.Set("Authorization", "token "+ghosthubKey)
	res := utility.TryResponse(client.Do(req))

	dataString := utility.ReadResponse(res)

	//TODO: just a test value, replace with handler
	tVal := utility.StringToJsonArray(dataString)

	fmt.Println("Users: ")
	for _, val := range tVal {
		fmt.Println(val["ldap_dn"])
	}
}

//GetGithubFileContent -> Get file content from github given a filename
func GetGithubFileContent(fileName string) string {
	response, err := http.Get(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1) //TODO: remove and replace with goroutine exit
	}

	dataString := utility.ReadResponse(response)

	return utility.DecodeReadResponse(dataString, "content")
}
