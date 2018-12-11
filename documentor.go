package main

import (
	gitscraper "documentor/service"
)

func main() {
	gitscraper.GetOrganizationMembers("PlayerBehavior")
}
