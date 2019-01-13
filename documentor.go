package main

import (
	gitscraper "github.com/mhandria/documentor/service"
)

func main() {
	gitscraper.GetOrganizationMembers("PlayerBehavior")
}
