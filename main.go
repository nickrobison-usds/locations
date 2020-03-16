package main

import (
	"fmt"
	"github.com/nickrobison-usds/test-locations/responses"
	"log"
	"net/http"
	"os"
)

func handle(locations *LocationList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "We have %d locations.\n", len(locations.locations))
	}
}

func main() {

	locations := NewLocationList(100)

	// Update the location list
	updateLocations(locations)
	// Run the server
	http.Handle("/", handle(locations))
	log.Println("Listening")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func updateLocations(locations *LocationList) {
	fmt.Println("Downloading the latest responses and formatting them")

	r, err := responses.New(os.Getenv("SHEET_ID"), os.Getenv("CREDENTIALS"))
	if err != nil {
		panic(err)
	}

	resp, err := r.GetResponses()
	if err != nil {
		panic(err)
	}

	log.Printf("Downloaded: %d\n", len(resp))

	for _, loc := range resp {
		locations.AddLocation(loc)
	}
}
