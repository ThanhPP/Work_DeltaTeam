package rebrandly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
	ggs "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/googlesheett"
)

func getDataFromEnv() (apiKey string, domainID string) {
	var err error
	apiKey, err = config.GetEnvKey("REBRANDLYAPIKEY")
	if err != nil {
		log.Println("getDataFromEnv-apiKey : ", err)
		return "", ""
	}

	domainID, err = config.GetEnvKey("REBRANDLYDOMAINID")
	if err != nil {
		log.Println("getDataFromEnv-domainID : ", err)
		return "", ""
	}

	return apiKey, domainID
}

//ShortLinkByRebrand
func shortLinkByRebrand(forwardLinkSlice []string, slashTagSlice []string) (shortLinkResult []string, successCount int, errorCount int) {
	apiKey, domainID := getDataFromEnv()
	for i := 0; i < len(forwardLinkSlice); i++ {
		resp, err := http.Get("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
		// fmt.Println("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
		if err != nil {
			log.Println("Err shortLinkByRebrand")
		}
		defer resp.Body.Close()

		fmt.Println(forwardLinkSlice[i]+" => https://rebrand.ly/"+slashTagSlice[i], " : ", resp.StatusCode)
		if resp.StatusCode == 200 {
			shortLinkResult = append(shortLinkResult, "https://rebrand.ly/"+slashTagSlice[i])
			successCount++
		} else {
			shortLinkResult = append(shortLinkResult, "error : https://rebrand.ly/"+slashTagSlice[i])
			errorCount++
		}
	}
	return shortLinkResult, successCount, errorCount
}

type linkCountType struct {
	Count int `json:"count"`
}

func countLinkRebranly() int {
	apikey, _ := getDataFromEnv()
	//START : read number of shortlink created
	req, err := http.NewRequest("GET", "https://api.rebrandly.com/v1/links/count", nil)
	if err != nil {
		log.Println("countLinkRebranly")
		return -1
	}
	req.Header.Set("Apikey", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("countLinkRebranly")
		return -1
	}
	defer resp.Body.Close()

	countByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("countLinkRebranly")
		return -1
	}

	var linkCounter linkCountType
	err = json.Unmarshal(countByte, &linkCounter)
	if err != nil {
		log.Println("countLinkRebranly")
		return -1
	}
	//END : read number of shortlink created

	return linkCounter.Count
}

//CreateShortLinkRebrandly get data from gg sheet and using rebrandly to create shortLink
func CreateShortLinkRebrandly(inputRange string, inputFwdLinks []string) (shortLinkResult []string, successCount int, errorCount int, usedCount int) {
	//Get range
	firstNum, secondNum, err := ggs.ParseRange(inputRange)
	if err != nil {
		log.Printf("\n\t ForwardLinks : %+v \n", err)
	}
	//Column assign
	slashTagCol := "W"

	//CreateShortLink
	slashTagSlice := ggs.GetDataFromRage(ggs.NewRange(firstNum, secondNum, slashTagCol))
	shortLinkResult, successCount, errorCount = shortLinkByRebrand(inputFwdLinks, slashTagSlice)

	usedCount = countLinkRebranly()
	return shortLinkResult, successCount, errorCount, usedCount
}