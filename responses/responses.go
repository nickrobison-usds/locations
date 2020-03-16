package responses

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"strings"
)

type LocationResponse struct {
	ID           int
	Timestamp    string
	Name         string
	Address      string
	City         string
	State        string
	Zip          string
	Days         []string
	StartTime    string
	EndTime      string
	TestDate     string
	Tests        string
	ContactName  string
	ContactPhone string
	ContactEmail string
	NewLocation  bool
	TimeZone     string
}

type ResponseProcessor struct {
	sheetID string
	service *sheets.SpreadsheetsService
}

func New(sheetID string, key string) (*ResponseProcessor, error) {
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsFile(key))
	if err != nil {
		return nil, err
	}

	s := sheets.NewSpreadsheetsService(srv)

	return &ResponseProcessor{sheetID: sheetID, service: s}, nil
}

func (r *ResponseProcessor) GetResponses() ([]LocationResponse, error) {
	readRange := "Form Responses 1!A2:U"
	resp, err := r.service.Values.Get(r.sheetID, readRange).Do()
	if err != nil {
		return nil, err
	}

	locations := make([]LocationResponse, 0)

	for idx, row := range resp.Values {
		fmt.Printf("Response: %s\n", row[0])

		timestamp := row[0].(string)
		//fmt.Println(timestamp)
		//if timestamp == "" {
		//	break
		//}

		// Parse a LocationResponse and add it to the array.
		// The Google Sheet is pretty messed up, so we have to skip around a bit.
		locations = append(locations, LocationResponse{
			ID:           idx,
			Timestamp:    timestamp,
			Name:         row[2].(string),
			Address:      row[3].(string),
			City:         row[4].(string),
			State:        row[5].(string),
			Zip:          row[6].(string),
			Days:         strings.Split(row[7].(string), ","),
			StartTime:    row[8].(string),
			EndTime:      row[9].(string),
			TestDate:     row[13].(string),
			Tests:        row[14].(string),
			ContactName:  row[15].(string),
			ContactPhone: row[16].(string),
			ContactEmail: row[17].(string),
			NewLocation:  row[18].(string) == "Yes",
			TimeZone:     row[19].(string),
		})
	}

	return locations, nil

}
