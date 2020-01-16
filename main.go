package main

var page string

type HydraShop struct {
	Title  string
	Text   string
	Market string
	Price  string
}

var hs = []HydraShop{}

func main() {
	go startServer()
	startBot()
}
