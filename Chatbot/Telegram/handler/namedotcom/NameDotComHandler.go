package namedotcom

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
	ggs "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/googlesheett"
)

var (
	nameConfig config.NameDotConfig
)

//GetDataFromEnv get default data from env file
func getDataFromEnv() {
	config.GetNameConfigObj()
}

//createForwardLinks receive store links and tempForwardlink to create forward link via name.com
func createForwardLinks(storeLinks []string, tempForwardLinks []string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
	getDataFromEnv()
	nameDomain := nameConfig.Domain
	nameAPIKey := nameConfig.APIKey
	nameUsername := nameConfig.Username

	for i := 0; i < len(tempForwardLinks); i++ {
		body := strings.NewReader(`{"host":"` + tempForwardLinks[i] + `.` + nameDomain + `","forwardsTo":"` + storeLinks[i] + `","type":"redirect"}`)

		req, err := http.NewRequest("POST", "https://api.name.com/v4/domains/"+nameDomain+"/url/forwarding", body)
		if err != nil {
			log.Println("CreateForwardLinks : ", err)
			forwardResult = append(forwardResult, fmt.Sprint("CreateForwardLinks - NewRequest - err"))
			continue
		}
		req.SetBasicAuth(nameUsername, nameAPIKey)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("CreateForwardLinks : ", err)
			forwardResult = append(forwardResult, fmt.Sprint("CreateForwardLinks - MakeRequest - err"))
			continue
		}
		resp.Body.Close()

		if resp.StatusCode == 200 {
			forwardResult = append(forwardResult, tempForwardLinks[i]+"."+nameDomain)
			successForwardCount++
		} else {
			forwardResult = append(forwardResult, "error : "+tempForwardLinks[i]+"."+nameDomain)
			errorForwardCount++
		}
	}

	return forwardResult, successForwardCount, errorForwardCount
}

//ForwardLinks create forward links from input range
func ForwardLinks(inputRange string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
	getDataFromEnv()
	//Get range
	firstNum, secondNum, err := ggs.ParseRange(inputRange)
	if err != nil {
		log.Printf("\n\t ForwardLinks : %+v \n", err)
		return nil, -1, -1
	}
	//Column assign
	storeLinksCol := nameConfig.StoreLinkColumn
	tempForwardLinksCol := nameConfig.TempForwardLinkColumn

	if dataRange := (secondNum - firstNum); dataRange <= 2 {
		storeLinks := ggs.GetDataFromRage(ggs.NewRange(firstNum, secondNum, storeLinksCol))
		tempForwardLinks := ggs.GetDataFromRage(ggs.NewRange(firstNum, secondNum, tempForwardLinksCol))
		forwardResult, successForwardCount, errorForwardCount = createForwardLinks(storeLinks, tempForwardLinks)

		return forwardResult, successForwardCount, errorForwardCount
	}

	//CONCURRENCY
	//Variable
	var wg sync.WaitGroup

	var forwardResult1 []string
	var successForwardCount1 int
	var errorForwardCount1 int

	var forwardResult2 []string
	var successForwardCount2 int
	var errorForwardCount2 int

	//Concurrent func
	wg.Add(2)

	go func() {
		//fmt.Println(firstNum, int(secondNum-((secondNum-firstNum)/2)))
		storeLinks1 := ggs.GetDataFromRage(ggs.NewRange(firstNum, int(secondNum-((secondNum-firstNum)/2)), storeLinksCol))
		tempForwardLinks1 := ggs.GetDataFromRage(ggs.NewRange(firstNum, int(secondNum-((secondNum-firstNum)/2)), tempForwardLinksCol))

		forwardResult1, successForwardCount1, errorForwardCount1 = createForwardLinks(storeLinks1, tempForwardLinks1)

		wg.Done()
	}()

	go func() {
		//fmt.Println(int(secondNum-((secondNum-firstNum)/2)+1), secondNum)
		storeLinks2 := ggs.GetDataFromRage(ggs.NewRange(int(secondNum-((secondNum-firstNum)/2)+1), secondNum, storeLinksCol))
		tempForwardLinks2 := ggs.GetDataFromRage(ggs.NewRange(int(secondNum-((secondNum-firstNum)/2)+1), secondNum, tempForwardLinksCol))

		forwardResult2, successForwardCount2, errorForwardCount2 = createForwardLinks(storeLinks2, tempForwardLinks2)

		wg.Done()
	}()

	wg.Wait()

	//Main phase
	forwardResult = append(forwardResult1, forwardResult2...)
	successForwardCount = successForwardCount1 + successForwardCount2
	errorForwardCount = errorForwardCount1 + errorForwardCount2

	return forwardResult, successForwardCount, errorForwardCount
}
