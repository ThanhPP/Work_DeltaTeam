package rebrandly

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
	ggs "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/googlesheett"
)

var (
	rbConfig *config.RBConfig
)

func getDataFromEnv() {
	rbConfig = config.GetRBConfigObj()
}

//ShortLinkByRebrand
func shortLinkByRebrand(forwardLinkSlice []string, slashTagSlice []string) (shortLinkResult []string, successCount int, errorCount int) {
	getDataFromEnv()
	apiKey := rbConfig.APIKey
	domainID := rbConfig.DomainID

	if len(forwardLinkSlice) != len(slashTagSlice) {
		log.Printf("\n\n\tshortLinkByRebrand : forwardLinkSlice-%+v slashTagSlice-%+v\n\n", len(forwardLinkSlice), len(slashTagSlice))
		shortLinkResult = append(shortLinkResult, "Input length not match")
		return shortLinkResult, -1, -1
	}

	for i := 0; i < len(forwardLinkSlice); i++ {
		resp, err := http.Get("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
		// fmt.Println("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
		if err != nil {
			log.Println("Err shortLinkByRebrand")
			shortLinkResult = append(shortLinkResult, "shortLinkByRebrand GET REQUEST ERROR")
			continue
		}
		defer resp.Body.Close()
		//fmt.Println(forwardLinkSlice[i]+" => https://rebrand.ly/"+slashTagSlice[i], " : ", resp.StatusCode)
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
	getDataFromEnv()
	apikey := rbConfig.APIKey
	//START : read number of shortlink created
	req, err := http.NewRequest("GET", "https://api.rebrandly.com/v1/links/count", nil)
	if err != nil {
		log.Println("countLinkRebranly error")
		return -1
	}
	req.Header.Set("Apikey", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("countLinkRebranly error")
		return -1
	}
	defer resp.Body.Close()

	countByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("countLinkRebranly error")
		return -1
	}

	var linkCounter linkCountType
	err = json.Unmarshal(countByte, &linkCounter)
	if err != nil {
		log.Println("countLinkRebranly error")
		return -1
	}
	//END : read number of shortlink created

	return linkCounter.Count
}

//CreateShortLinkRebrandly get data from gg sheet and using rebrandly to create shortLink
func CreateShortLinkRebrandly(inputRange string, inputFwdLinks []string) (shortLinkResult []string, successCount int, errorCount int, usedCount int) {
	getDataFromEnv()
	//Get range
	firstNum, secondNum, err := ggs.ParseRange(inputRange)
	if err != nil {
		log.Printf("\n\t ForwardLinks : %+v \n", err)
		return nil, -1, -1, -1
	}
	//Column assign
	slashTagCol := rbConfig.SlashTagColumn

	//CreateShortLink
	if dataRange := (secondNum - firstNum); dataRange <= 2 {
		slashTagSlice := ggs.GetDataFromRage(ggs.NewRange(firstNum, secondNum, slashTagCol))
		shortLinkResult, successCount, errorCount = shortLinkByRebrand(inputFwdLinks, slashTagSlice)

		usedCount = countLinkRebranly()

		return shortLinkResult, successCount, errorCount, usedCount
	}

	//CONCURRENCY
	//Variable
	var wg sync.WaitGroup

	var shortLinkResult1 []string
	var successCount1 int
	var errorCount1 int

	var shortLinkResult2 []string
	var successCount2 int
	var errorCount2 int

	//Concurrent func
	wg.Add(2)

	go func() {
		slashTagSlice1 := ggs.GetDataFromRage(ggs.NewRange(firstNum, int(secondNum-((secondNum-firstNum)/2))-1, slashTagCol))
		shortLinkResult1, successCount1, errorCount1 = shortLinkByRebrand(inputFwdLinks[:len(inputFwdLinks)/2], slashTagSlice1)

		wg.Done()
	}()

	go func() {
		slashTagSlice2 := ggs.GetDataFromRage(ggs.NewRange(int(secondNum-((secondNum-firstNum)/2)), secondNum, slashTagCol))
		shortLinkResult2, successCount2, errorCount2 = shortLinkByRebrand(inputFwdLinks[len(inputFwdLinks)/2:], slashTagSlice2)

		wg.Done()
	}()

	wg.Wait()

	// Return result
	shortLinkResult = append(shortLinkResult1, shortLinkResult2...)
	successCount = successCount1 + successCount2
	errorCount = errorCount1 + errorCount2

	usedCount = countLinkRebranly()
	return shortLinkResult, successCount, errorCount, usedCount
}

func checkAPIKey(apikey string) error {
	//Use count link request to check if apikey is valid or not
	req, err := http.NewRequest("GET", "https://api.rebrandly.com/v1/links/count", nil)
	if err != nil {
		log.Println("countLinkRebranly req error")
		return err
	}
	req.Header.Set("Apikey", apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("countLinkRebranly resp error")
		return err
	}
	defer resp.Body.Close()

	countByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("countLinkRebranly countByte error")
		return err
	}

	var linkCounter linkCountType
	err = json.Unmarshal(countByte, &linkCounter)
	if err != nil {
		log.Println("countLinkRebranly Unmarshal error")
		return err
	}

	if linkCounter.Count == -1 {
		return errors.New("Not an API key for rebrandly")
	}

	return nil
}

//SetRebrandlyAPIKey set new rebrandly to the config file
func SetRebrandlyAPIKey(apikey string) error {
	getDataFromEnv()

	err := checkAPIKey(apikey)
	if err != nil {
		log.Printf("Can not validate the api key : %+v", err)
		return err
	}

	err = rbConfig.ChangeAPIKey(apikey)
	if err != nil {
		log.Printf("Can Change API Key : %+v", err)
		return err
	}

	return nil
}
