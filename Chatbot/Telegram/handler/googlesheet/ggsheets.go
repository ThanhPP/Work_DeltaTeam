package googlesheet

// https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/get?apix_params=%7B%22spreadsheetId%22%3A%221alTTN3sIuK-d-M16lV6gPK4M6AwBkc7UQj2nA5ZBol0%22%2C%22range%22%3A%22T3653%3AT3663%22%7D

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

//CreateClient create a google sheet service using pre-config
func CreateClient() (*sheets.Service, error) {
	b, err := ioutil.ReadFile("/home/tpp18/go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/googlesheet/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv, nil
}

//GetDataFromRage get data from given range and sheetID
func GetDataFromRage(sheetRange string) {
	// constant
	spreadsheetID := "1alTTN3sIuK-d-M16lV6gPK4M6AwBkc7UQj2nA5ZBol0"
	readRange := "Product!" + sheetRange
	//create a service
	srv, err := CreateClient()
	if err != nil {
		log.Fatalf("Unable to create a service : %+v \n", err)
	}
	//fetch data
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v \n", err)
	}
	if len(resp.Values) == 0 {
		fmt.Println("found nothing")
	} else {
		for _, row := range resp.Values {
			fmt.Println(row)
		}
	}

}
