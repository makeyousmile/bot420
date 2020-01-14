package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func parse(page string) string {
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
		hsSlise = append(hsSlise, hs)
		log.Print(hs.Title)
		log.Print(hs.Text)
		log.Print(hs.Market)
		log.Print(hs.Price)
	})

	return page
}
