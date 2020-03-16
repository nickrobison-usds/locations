package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
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

	// Update the list every couple of minutes
	interval, err := strconv.ParseInt(os.Getenv("INTERVAL"), 10, 32)
	if err != nil {
		panic(err)
	}
	log.Printf("Updating every %d minutes\n", interval)
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	go func(locations *LocationList) {
		for {
			updateLocations(locations)
			<-ticker.C
		}
	}(locations)

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
