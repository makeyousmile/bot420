package main

import "time"

var page string

type HydraShop struct {
	Title      string
	Text       string
	Market     string
	Price      string
	UpdateTime time.Time
}

var hs = []HydraShop{}

func main() {
	go startServer()
	startBot()
}
