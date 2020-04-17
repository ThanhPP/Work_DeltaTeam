package namedotcom

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

	//Parse range
	storeLinks := ggs.GetDataFromRage(ggs.NewRange(firstNum, secondNum, storeLinksCol))
	tempForwardLinks := ggs.GetDataFromRage(ggs.NewRange(firstNum, secondNum, tempForwardLinksCol))

	//Main phase
	forwardResult, successForwardCount, errorForwardCount = createForwardLinks(storeLinks, tempForwardLinks)

	return forwardResult, successForwardCount, errorForwardCount
}
