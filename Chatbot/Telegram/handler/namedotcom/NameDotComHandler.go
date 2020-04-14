package namedotcom

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
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

//CreateForwardLinks receive store links and tempForwardlink to create forward link via name.com
func CreateForwardLinks(storeLinks []string, tempForwardLinks []string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
	nameDomain, nameAPIKey, nameUsername := GetDataFromEnv()
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
