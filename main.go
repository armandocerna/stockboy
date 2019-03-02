package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"log"
	"os"
	"time"
)

var (
	desiredStock string
	//strikePrice float64
	currentPrice float64
)

func main() {
	desiredStock := os.Args[1]
	now := time.Now()

	today := now.Format("2006-01-02")
	tomorrow := now.Add(time.Hour * 24).Format("2006-01-02")

	fmt.Println(desiredStock, today, tomorrow)

	q, err := quote.NewQuoteFromYahoo(desiredStock, today, tomorrow, quote.Daily, true)
	if err != nil {
		log.Fatalf("error retrieving quote %v", err)
	}
	fmt.Println("Quote: %s", q)
}