package main

import (
	"github.com/nickrobison-usds/test-locations/responses"
	"sync"
)

const Template = `
<head>
</head>
<body>
<h2>Co-Vid Testing Locations</h2>
{{range .Locations}}
<div itemscope itemtype="http://schema.org/Place">
	<div itemscope="" itemprop="identifier" itemptype="http://schema.org/PropertyValue">
	<h4><span itemprop="propertyID">Location Number:</span>: <span itemprop="value">{{.ID}}</span></h4>
	</div>
	Name: <span itemprop="name">{{.Name}}</span>
	<div itemprop="address" itemscope itemtype="http://schema.org/PostalAddress">
		Address:
		<span itemprop="streetAddress">{{.Address}}</span>
		<span itemprop="addressLocality">{{.City}}</span>
		<span itemprop="addressRegion">{{.State}}</span>
		<span itemprop="postalCode">{{.Zip}}</span>
	</div>
	<div>
	Hours of Operation:
	{{ $save := . }}
	{{range .Days}}
		<div itemprop="openingHoursSpecification" itemscope itemtype="http://schema.org/OpeningHoursSpecification">
        <link itemprop="dayOfWeek" href="http://schema.org/{{.}}">{{.}}: <time itemprop="opens" content="{{$save.StartTime}}">{{$save.StartTime}}</time> - <time itemprop="closes" content="{{$save.EndTime}}">{{$save.EndTime}}</time>
    </div>
	{{end}}
	</div>
</div>
{{end}}
</body>`

// LocationList is safe to use concurrently
type LocationList struct {
	Locations []responses.LocationResponse
	mux       sync.Mutex
}

func (l *LocationList) AddLocation(loc responses.LocationResponse) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.Locations = append(l.Locations, loc)
}

func NewLocationList() *LocationList {
	locations := make([]responses.LocationResponse, 0)
	return &LocationList{
		Locations: locations,
		mux:       sync.Mutex{},
	}
}
