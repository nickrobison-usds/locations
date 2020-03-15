package responses

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type LocationResponse struct {
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

func (r *ResponseProcessor) GetResponses() (string, error) {
	readRange := "Form Responses 1!A2:P"
	resp, err := r.service.Values.Get(r.sheetID, readRange).Do()
	if err != nil {
		return "", err
	}

	for _, row := range resp.Values {
		fmt.Printf("Response: %s\n", row[0])
	}

	return "", nil

}
