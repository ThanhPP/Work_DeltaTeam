package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	//Edit your PATH here
	"github.com/THANHPP/Work_DeltaTeam/Work_ShortLinkTool/config"
	"github.com/THANHPP/Work_DeltaTeam/Work_ShortLinkTool/handler"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkLen(slice1 []string, slice2 []string) {
	fmt.Println(len(slice1), "  ", len(slice2))
	if len(slice1) != len(slice2) {
		log.Fatalln("Len between file doesn't match")
	}
	fmt.Println("Check len : OK")
	fmt.Println()
}

func main() {
	//flag
	forward := flag.Bool("forward", false, "Boolean value to start forward phase")
	shortLink := flag.Bool("shortLink", false, "Boolean value to start short link phase")
	flag.Parse()

	fmt.Println()
	fmt.Println("forward : ", *forward)
	fmt.Println("short link : ", *shortLink)
	fmt.Println()

	if *forward {

		// //Read files
		storeLinks, err := handler.ReadFile(config.StoreLinkPath)
		checkErr(err)
		storeLinks1 := storeLinks[:len(storeLinks)/2]
		storeLinks2 := storeLinks[len(storeLinks)/2:]

		tempForwardLinks, err := handler.ReadFile(config.TempForwardLinkPath)
		checkErr(err)
		tempForwardLinks1 := tempForwardLinks[:len(tempForwardLinks)/2]
		tempForwardLinks2 := tempForwardLinks[len(tempForwardLinks)/2:]

		checkLen(storeLinks, tempForwardLinks)

		//Forwarding using name.com
		fmt.Println("----------START FORWARD PHASE----------")
		timerForwardStart := time.Now()

		var wgForward sync.WaitGroup
		//func1
		var result1 []string
		var sucCount1 int
		var errCount1 int
		//func2
		var result2 []string
		var sucCount2 int
		var errCount2 int

		wgForward.Add(2)
		go func() {
			defer wgForward.Done()
			result1, sucCount1, errCount1 = handler.CreateForwardLink(storeLinks1, tempForwardLinks1, config.NameAPIKey1)
		}()
		go func() {
			defer wgForward.Done()
			result2, sucCount2, errCount2 = handler.CreateForwardLink(storeLinks2, tempForwardLinks2, config.NameAPIKey1)
		}()
		wgForward.Wait()

		result := append(result1, result2...)

		fmt.Println()
		fmt.Println("SUCCESS FORWARDLINK CREATED : ", sucCount1+sucCount2)
		fmt.Println("ERROR FORWARDLINK CREATED : ", errCount1+errCount2)
		fmt.Println("forward create in time : ", time.Since(timerForwardStart))
		fmt.Println()
		//show forward result
		f, err := os.Create(config.TempResultForwardLinkPath)
		checkErr(err)
		for _, forwardLinkCreated := range result[:len(result)-1] {
			fmt.Println(forwardLinkCreated)
			f.WriteString(forwardLinkCreated + "\n")
		}
		fmt.Println(result[len(result)-1])
		f.WriteString(result[len(result)-1])
		f.Close()
		fmt.Println()
		fmt.Println("----------FORWARD PHASE END HERE----------")
		fmt.Println()
		//End Forwarding using name.com
	}

	if *shortLink {
		//Shortlink using rebrand
		fmt.Println("----------SHORTLINK PHASE START HERE----------")
		timerShortLinkStart := time.Now()

		//Read files
		tempForwardLinks, err := handler.ReadFile(config.TempForwardLinkPath)
		slashTag, err := handler.ReadFile(config.SlashTagPath)
		checkErr(err)
		slashTag1 := slashTag[:len(slashTag)/2]
		slashTag2 := slashTag[len(slashTag)/2:]
		checkLen(slashTag, tempForwardLinks)

		forwardLinkSlice, err := handler.ReadFile(config.TempResultForwardLinkPath)
		checkErr(err)
		checkLen(forwardLinkSlice, slashTag)
		forwardLinkSlice1 := forwardLinkSlice[:len(forwardLinkSlice)/2]
		forwardLinkSlice2 := forwardLinkSlice[len(forwardLinkSlice)/2:]

		//variables
		var wgShortlink sync.WaitGroup
		//func1
		var shortLinkResult1 []string
		var shortLinkSucCount1 int
		var shortLinkErrCount1 int
		//func2
		var shortLinkResult2 []string
		var shortLinkSucCount2 int
		var shortLinkErrCount2 int

		wgShortlink.Add(2)
		go func() {
			shortLinkResult1, shortLinkSucCount1, shortLinkErrCount1 = handler.ShortLinkByRebrand(forwardLinkSlice1, slashTag1, config.RebrandlyAPIKey, config.RebrandlyDomainID)
			wgShortlink.Done()
		}()
		go func() {
			shortLinkResult2, shortLinkSucCount2, shortLinkErrCount2 = handler.ShortLinkByRebrand(forwardLinkSlice2, slashTag2, config.RebrandlyAPIKey, config.RebrandlyDomainID)
			wgShortlink.Done()
		}()
		wgShortlink.Wait()

		fmt.Println()
		fmt.Println("SUCCESS SHORT LINK CREATED : ", shortLinkSucCount1+shortLinkSucCount2)
		fmt.Println("ERROR SHORT LINK CREATED : ", shortLinkErrCount1+shortLinkErrCount2)
		fmt.Println("forward create in time : ", time.Since(timerShortLinkStart))
		fmt.Println()

		shortLinkResult := append(shortLinkResult1, shortLinkResult2...)
		for _, shortLinkCreated := range shortLinkResult {
			fmt.Println(shortLinkCreated)
		}
		fmt.Println("----------SHORTLINK PHASE END HERE----------")
		fmt.Println()

		//Check limit
		linkCounter := handler.CountLinkRebranly(config.RebrandlyAPIKey)
		fmt.Println("Created ", linkCounter, " links with this API KEY")
		fmt.Println("Rebrandly ", 500-linkCounter, " links left")
		//End Shortlink using rebrand
	}

	fmt.Printf("\n PROGRAM END HERE \n")
}
