package main

import (
	gitscraper "ghosthub.com/mhandria/documentor/service"
)

func main() {
	gitscraper.GetOrganizationMembers("PlayerBehavior")
}
