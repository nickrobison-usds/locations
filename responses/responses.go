package responses

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type LocationResponse struct {
	Timestamp  string
	Email      string
	Name       string
	Address    string
	City       string
	State      string
	Zip        string
	Days       string
	StartTime  string
	EndTime    string
	Phone      string
	Tests      int
	TaxID      string
	FEMA       bool
	LocationID string
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
	readRange := "Form Responses 1!A2:P"
	resp, err := r.service.Values.Get(r.sheetID, readRange).Do()
	if err != nil {
		return nil, err
	}

	locations := make([]LocationResponse, 10)

	for _, row := range resp.Values {
		fmt.Printf("Response: %s\n", row[0])

		locations = append(locations, LocationResponse{
			Timestamp:  row[0].(string),
			Email:      "",
			Name:       "",
			Address:    "",
			City:       "",
			State:      "",
			Zip:        "",
			Days:       "",
			StartTime:  "",
			EndTime:    "",
			Phone:      "",
			Tests:      0,
			TaxID:      "",
			FEMA:       false,
			LocationID: "",
		})
	}

	return locations, nil

}
