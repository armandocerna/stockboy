package main

import (
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
	desiredStock := os.Args[0]

	today, err := time.Parse("2006-01-02", time.Now())
	if err != nil {
		log.Fatalf("unable to parse time object: %v", err)
	}

	tomorrow, err := time.Parse("2006-01-02", (time.Now() + time.Hour*24))
	if err != nil {
		log.Fatalf("unable to parse time object: %v", err)
	}

	quote.NewQuoteFromYahoo(desiredStock, today, tomorrow)
}