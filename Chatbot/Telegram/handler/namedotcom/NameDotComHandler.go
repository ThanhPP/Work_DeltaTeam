package namedotcom

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func parseRange(inputRange string) (int, int, error) {
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

func newRange(firstNum int, secondNum int, col string) string {
	newStr := col + strconv.Itoa(firstNum) + ":" + col + strconv.Itoa(secondNum)
	return newStr
}

//ForwardLinks create forward links from input range
func ForwardLinks(inputRange string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
	//Get range
	firstNum, secondNum, err := parseRange(inputRange)
	if err != nil {
		log.Printf("\n\t ForwardLinks : %+v \n", err)
	}

	//Column assign
	storeLinksCol := "T"
	tempForwardLinksCol := "U"

	//Parse range
	storeLinks := ggs.GetDataFromRage(newRange(firstNum, secondNum, storeLinksCol))
	tempForwardLinks := ggs.GetDataFromRage(newRange(firstNum, secondNum, tempForwardLinksCol))

	//Main phase
	forwardResult, successForwardCount, errorForwardCount = createForwardLinks(storeLinks, tempForwardLinks)

	return forwardResult, successForwardCount, errorForwardCount
}
