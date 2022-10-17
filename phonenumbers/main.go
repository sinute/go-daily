package main

import (
	"github.com/nyaruka/phonenumbers"
)

func main() {
	// parse our phone number
	num, err := phonenumbers.Parse("11067800", "CN")
	if err != nil {
		panic(err)
	}

	// format it using national format
	formattedNum := phonenumbers.Format(num, phonenumbers.E164)
	println(formattedNum)
}
