package main

import (
	"flag"
	"github.com/gideon-maina/rss-processor/fetchrss"
	"github.com/gideon-maina/rss-processor/serverss"
	"log"
	"time"
)

func main() {
	log.Println("Starting RSS Processor...")
	var refreshInterval int
	flag.IntVar(&refreshInterval, "refresh", 5, "The refresh interval to fetch feeds in minutes.")
	flag.Parse()

	log.Println("RSS processor will fetch and update feeds every", refreshInterval, " minutes.")
	// Open a server to serve client requests here in it's own routine
	go serverss.ServeClients()
	// First time being called fetch feeds
	fetchrss.FetchAndStoreRSSFeeds()
	// Periodically run the fetch and store RSS feeds function every supplied refresh minutes default 5
	updateTimer := time.NewTicker(time.Duration(refreshInterval) * time.Minute)
	for {
		select {
		case <-updateTimer.C:
			fetchrss.FetchAndStoreRSSFeeds()
		}
	}
}
