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

//GetDataFromEnv get default data from env file
func getDataFromEnv() (nameDomain string, nameAPIKey string, nameUsername string) {
	var err error
	nameDomain, err = config.GetEnvKey("NAMEDOMAIN")
	if err != nil {
		log.Println("getDataFromEnv : ", err)
		return "", "", ""
	}
	fmt.Println(nameDomain)

	nameAPIKey, err = config.GetEnvKey("NAMEAPIKEY")
	if err != nil {
		log.Println("getDataFromEnv : ", err)
		return "", "", ""
	}

	nameUsername, err = config.GetEnvKey("NAMEUSRNAME")
	if err != nil {
		log.Println("getDataFromEnv : ", err)
		return "", "", ""
	}
	return nameDomain, nameAPIKey, nameUsername
}

//createForwardLinks receive store links and tempForwardlink to create forward link via name.com
func createForwardLinks(storeLinks []string, tempForwardLinks []string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
	nameDomain, nameAPIKey, nameUsername := getDataFromEnv()
	for i := 0; i < len(tempForwardLinks); i++ {
		body := strings.NewReader(`{"host":"` + tempForwardLinks[i] + `.` + nameDomain + `","forwardsTo":"` + storeLinks[i] + `","type":"redirect"}`)

		req, err := http.NewRequest("POST", "https://api.name.com/v4/domains/"+nameDomain+"/url/forwarding", body)
		if err != nil {
			log.Println("CreateForwardLinks : ", err)
		}
		req.SetBasicAuth(nameUsername, nameAPIKey)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("CreateForwardLinks : ", err)
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
	//Get range
	firstNum, secondNum, err := ggs.ParseRange(inputRange)
	if err != nil {
		log.Printf("\n\t ForwardLinks : %+v \n", err)
	}
	//Column assign
	storeLinksCol := "T"
	tempForwardLinksCol := "U"

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
		fmt.Println(firstNum, int(secondNum-((secondNum-firstNum)/2)))
		storeLinks1 := ggs.GetDataFromRage(ggs.NewRange(firstNum, int(secondNum-((secondNum-firstNum)/2)), storeLinksCol))
		tempForwardLinks1 := ggs.GetDataFromRage(ggs.NewRange(firstNum, int(secondNum-((secondNum-firstNum)/2)), tempForwardLinksCol))

		forwardResult1, successForwardCount1, errorForwardCount1 = createForwardLinks(storeLinks1, tempForwardLinks1)

		wg.Done()
	}()

	go func() {
		fmt.Println(int(secondNum-((secondNum-firstNum)/2)+1), secondNum)
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
