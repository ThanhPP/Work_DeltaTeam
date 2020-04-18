package googlesheett

// https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/get?apix_params=%7B%22spreadsheetId%22%3A%221alTTN3sIuK-d-M16lV6gPK4M6AwBkc7UQj2nA5ZBol0%22%2C%22range%22%3A%22T3653%3AT3663%22%7D
// https://www.prudentdevs.club/gsheets-go

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
)

var (
	linuxGgsJSONPath, _   = config.GetEnvKey("LINUXGGSSCRPATH")
	windowsGgsJSONPath, _ = config.GetEnvKey("WINDOWSGGSSCRPATH")
)

//CreateClient create a google sheet service using pre-config
func createClient() (*sheets.Service, error) {
	var path string
	if runtime.GOOS == "windows" {
		path = windowsGgsJSONPath
	} else {
		path = linuxGgsJSONPath
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(context.TODO())

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv, nil
}

//GetDataFromRage get data from given range and sheetID
func GetDataFromRage(sheetRange string) (rows []string) {
	// constant
	spreadsheetID, err := config.GetEnvKey("SPREADSHEETID")
	if err != nil {
		log.Printf("GetDataFromRage : Error %+v \n", err)
	}
	readRange := "Product!" + sheetRange
	//create a service
	srv, err := createClient()
	if err != nil {
		log.Printf("GetDataFromRage :Unable to create a service : %+v \n", err)
	}
	//fetch data
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Printf("GetDataFromRage : Unable to retrieve data from sheet: %v \n", err)
	}
	if len(resp.Values) == 0 {
		log.Println("GetDataFromRage : found nothing")
	} else {
		for _, row := range resp.Values {
			rows = append(rows, row[0].(string))
			// log.Println(i, "  ", row[0])
		}
		return rows
	}
	return nil
}

//ParseRange return number received from input range
func ParseRange(inputRange string) (int, int, error) {
	inputRange = strings.TrimSpace(inputRange)

	separateIndex := strings.Index(inputRange, ":")
	if separateIndex == -1 {
		return -1, -1, errors.New("inputRange : Can not find : in inputRange")
	}
	firstNum, err1 := strconv.Atoi(inputRange[:separateIndex])
	secondNum, err2 := strconv.Atoi(inputRange[separateIndex+1:])
	if err1 != nil || err2 != nil {
		log.Printf("inputRange : \n\t %+v \n\t %+v", err1, err2)
		return -1, -1, errors.New("inputRange : Atoi")
	}

	if firstNum > secondNum {
		return -1, -1, errors.New("Invalid range input")
	}
	return firstNum, secondNum, nil
}

//NewRange create range with input number and col
func NewRange(firstNum int, secondNum int, col string) string {
	newStr := col + strconv.Itoa(firstNum) + ":" + col + strconv.Itoa(secondNum)
	return newStr
}
