package service

import (
	"documentor/secret"
	"documentor/utility"
	"fmt"
	"log"
	"net/http"
	"os"
)

const ghosthubKey = secret.GhosthubKey
const ghosthubApiUrl = secret.GhosthubApiUrl

var client = &http.Client{}

func GetOrganizationMembers(orgName string) {
	endpoint := ghosthubApiUrl + "/orgs/" + orgName + "/members"
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

func GetGithubFileContent(fileName string) string {
	response, err := http.Get(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1) //TODO: remove and replace with goroutine exit
	}

	dataString := utility.ReadResponse(response)

	return utility.DecodeReadResponse(dataString, "content")
}
