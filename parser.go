package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
)

func parse(page string) []HydraShop {
	hsSlise := []HydraShop{}
	selection, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		log.Print(err)
	}
	selection.Find("div.desc").Each(func(i int, selection *goquery.Selection) {
		hs := HydraShop{}
		hs.Title = strings.TrimSpace(selection.Find("div.title").Text())
		hs.Text = strings.TrimSpace(selection.Find("div.text").Text())
		hs.Market = strings.TrimSpace(selection.Find("div.market").Text())
		hs.Price = strings.TrimSpace(selection.Find("span.price").Text())
		hs.UpdateTime = time.Now()
		hsSlise = append(hsSlise, hs)
		log.Print(hs.Title)
		log.Print(hs.Text)
		log.Print(hs.Market)
		log.Print(hs.Price)
		selection.Find("div.slide").Each(func(n int, selection *goquery.Selection) {
			subhs := HydraShop{}
			subhs.Title = strings.TrimSpace(selection.Find("div.slide_title").Text())
			subhs.Price = strings.TrimSpace(selection.Find("span.slide_price").Text())
			subhs.Market = hs.Market
			subhs.UpdateTime = time.Now()
			hsSlise = append(hsSlise, subhs)
		})
	})

	return hsSlise
}
