package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/THANHPP/Work_Crawler/handler/shopbase"
	"github.com/THANHPP/Work_Crawler/handler/woo"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	//START Crawl Shopbase
	go func() {
		var shopBaseProductList []shopbase.Products

		shopBaseCrawlTimeStart := time.Now()
		shopBaseProductList = shopbase.Crawl("https://www.ellezz.com")
		shopBaseCrawlTime := time.Since(shopBaseCrawlTimeStart)
		// for i, shopBaseproduct := range shopBaseProductList {
		// 	fmt.Println(i, " - ", shopBaseproduct.ID)
		// }
		fmt.Println()
		fmt.Println("Crawl ", len(shopBaseProductList), " products of https://www.ellezz.com in ", shopBaseCrawlTime)
		fmt.Println()
		wg.Done()
	}()
	//END Crawl Shopbaseode

	//START Crawl Woo
	go func() {
		wooCrawlTimeStart := time.Now()
		wooProductIDList := woo.Crawl("https://raccoonrider.com")
		wooCrawlTime := time.Since(wooCrawlTimeStart)

		// for i := 0; i < len(wooProductIDList); i++ {
		// 	fmt.Println(wooProductIDList[i])
		// }
		fmt.Println()
		fmt.Println("Crawl ", len(wooProductIDList), " products of https://raccoonrider.com in ", wooCrawlTime)

		wg.Done()
	}()
	//END Crawl Woo

	wg.Wait()
}
