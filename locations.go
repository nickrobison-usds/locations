package main

import (
	"github.com/nickrobison-usds/test-locations/responses"
	"sync"
)

// LocationList is safe to use concurrently
type LocationList struct {
	locations []responses.LocationResponse
	mux       sync.Mutex
}

func (l *LocationList) AddLocation(loc responses.LocationResponse) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.locations = append(l.locations, loc)
}

func NewLocationList(size int) *LocationList {
	locations := make([]responses.LocationResponse, size)
	return &LocationList{
		locations: locations,
		mux:       sync.Mutex{},
	}
}
