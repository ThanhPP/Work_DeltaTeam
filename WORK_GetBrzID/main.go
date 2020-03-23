package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func getID(nameData string) {
	c := colly.NewCollector()

	c.OnRequest(func(rq *colly.Request) {
		log.Println("Visiting : ", rq.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong !!!!")
	})

	c.OnHTML("body > div.container > div:nth-child(2)", func(e *colly.HTMLElement) {
		e.ForEach("body > div.container > div > a > p", func(_ int, e1 *colly.HTMLElement) {
			data := e1.Text
			if strings.Contains(data, "CNPJ") {
				fmt.Println(data)
			}
		})
	})
	c.Visit("https://cnpj.biz/procura/" + nameData)

}

func lastString(ss []string) string {
	return ss[len(ss)-1]
}

func getLastNameBrz() (s string) {
	c := colly.NewCollector()

	c.OnRequest(func(rq *colly.Request) {
		//log.Println("Visiting : ", rq.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong !!!!")
	})

	c.OnHTML("body > div.container.index.no-padding > div.row.main > div.col-md-9.col-sm-9.col-xs-12.main-left > div > div.row.detail.no-margin.no-padding > div.row > div:nth-child(3) > div.col-sm-8.col-xs-6.right", func(e *colly.HTMLElement) {
		data := e.ChildAttr("strong > input", "value")
		s = lastString(strings.Fields(data))
		fmt.Println(s)
	})

	c.Visit("https://www.fakeaddressgenerator.com/World/Brazil_address_generator")

	return s
}

func main() {
	name := getLastNameBrz()
	if len(name) <= 1 {
		panic("nothing")
	}
	getID(name)
}
