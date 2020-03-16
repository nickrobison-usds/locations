package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nickrobison-usds/test-locations/responses"
)

func handle(locations *LocationList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("webpage").Parse(Template)
		if err != nil {
			panic(err)
		}
		locations.mux.Lock()
		defer locations.mux.Unlock()
		t.Execute(w, locations)
	}
}

func main() {

	locations := NewLocationList()

	// Update the list every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)
	go func(locations *LocationList) {
		select {
		case <-ticker.C:
			updateLocations(locations)
		}
	}(locations)

	// Update the location list
	updateLocations(locations)
	// Run the server
	http.Handle("/", handle(locations))
	port := os.Getenv("PORT")
	log.Printf("Listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func updateLocations(locations *LocationList) {
	fmt.Println("Downloading the latest responses and formatting them")

	r, err := responses.New(os.Getenv("SHEET_ID"), []byte(os.Getenv("CREDENTIALS")))
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
