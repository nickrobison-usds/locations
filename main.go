package main

import (
	"fmt"
	"github.com/nickrobison-usds/test-locations/responses"
	"os"
)

func main() {
	fmt.Println("Downloading the latest responses and formatting them")

	r, err := responses.New(os.Getenv("SHEET_ID"), os.Getenv("CREDENTIALS"))
	if err != nil {
		panic(err)
	}

	_, err = r.GetResponses()
	if err != nil {
		panic(err)
	}
}
