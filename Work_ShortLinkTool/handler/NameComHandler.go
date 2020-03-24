package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/THANHPP/Work_DeltaTeam/Work_ShortLinkTool/config"
)

//CreateForwardLink Get Data from slice and create forwardlink using name.com API token
func CreateForwardLink(storeLinks []string, tempForwardLinks []string, apiKey string) (forwardResult []string, successForwardCount int, errorForwardCount int) {
	//Start to create forward link using name.com
	for i := 0; i < len(tempForwardLinks); i++ {
		body := strings.NewReader(`{"host":"` + tempForwardLinks[i] + `.` + config.NameDomain + `","forwardsTo":"` + storeLinks[i] + `","type":"redirect"}`)
		// fmt.Printf("\n %+v\n %+v\n %+v\n", tempForwardLinks[i], config.NameDomain, storeLinks[i])
		req, err := http.NewRequest("POST", "https://api.name.com/v4/domains/"+config.NameDomain+"/url/forwarding", body)
		checkErr(err)
		req.SetBasicAuth(config.NameUsername, apiKey)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// fmt.Printf("\n%+v\n\n", req)

		resp, err := http.DefaultClient.Do(req)
		checkErr(err)
		resp.Body.Close()

		// fmt.Printf("\n %+v\n %+v\n %+v\n", storeLinks[i], tempForwardLinks[i], resp.StatusCode)
		if resp.StatusCode == 200 {
			forwardResult = append(forwardResult, tempForwardLinks[i]+"."+config.NameDomain)
			successForwardCount++
		} else {
			forwardResult = append(forwardResult, "error : "+tempForwardLinks[i]+"."+config.NameDomain)
			errorForwardCount++
		}

		fmt.Println(storeLinks[i], " => ", tempForwardLinks[i], " : ", resp.StatusCode)
	}

	return forwardResult, successForwardCount, errorForwardCount
	//END CREATE FORWARD LINK
}

//GetForwardList Get all forward link of a domain
func GetForwardList() {
	req, err := http.NewRequest("GET", "https://api.name.com/v4/domains/"+"deltapodvn.com"+"/url/forwarding?page=1", nil)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth(config.NameUsername, config.NameAPIKey1)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)
}
