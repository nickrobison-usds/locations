package main

import (
	"sync"

	"github.com/nickrobison-usds/test-locations/responses"
)

// Template is the HTML template used to generate the Location outputs
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
        <link itemprop="dayOfWeek" href="http://schema.org/{{.}}">{{.}}: <time itemprop="opens" content="{{$save.StartTime_Schema}}">{{$save.StartTime_Hum}}</time> - <time itemprop="closes" content="{{$save.EndTime_Schema}}">{{$save.EndTime_Hum}}</time>
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

// AddLocation appends a new LocationResponse to the existing list
func (l *LocationList) AddLocation(loc responses.LocationResponse) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.Locations = append(l.Locations, loc)
}

//NewLocationList creates a new struct for holding the threadsafe list
func NewLocationList() *LocationList {
	return &LocationList{
		Locations: make([]responses.LocationResponse, 0),
		mux:       sync.Mutex{},
	}
}
