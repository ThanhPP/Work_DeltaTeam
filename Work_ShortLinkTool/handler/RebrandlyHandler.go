package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/THANHPP/Work_DeltaTeam/Work_ShortLinkTool/config"
)

type linkCountType struct {
	Count int `json:"count"`
}

//CountLinkRebranly to prevent 500 link limit
func CountLinkRebranly(apikey string) int {
	//START : read number of shortlink created
	req, err := http.NewRequest("GET", "https://api.rebrandly.com/v1/links/count", nil)
	if err != nil {
		checkErr(err)
	}
	req.Header.Set("Apikey", config.RebrandlyAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		checkErr(err)
	}
	defer resp.Body.Close()
	countByte, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var linkCounter linkCountType
	err = json.Unmarshal(countByte, &linkCounter)
	checkErr(err)
	//END : read number of shortlink created

	return linkCounter.Count
}

//ShortLinkByRebrand Use rebrand to short link
func ShortLinkByRebrand(forwardLinkSlice []string, slashTagSlice []string, apiKey string, domainID string) (shortLinkResult []string, successCount int, errorCount int) {
	for i := 0; i < len(forwardLinkSlice); i++ {
		resp, err := http.Get("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
		// fmt.Println("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
		checkErr(err)
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
