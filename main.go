package main

import (
	"github.com/gideon-maina/rss-processor/fetchrss"
	"github.com/gideon-maina/rss-processor/serverss"
	"log"
	"time"
)

func main() {
	log.Println("Starting RSS Processor...")
	// Open a server to serve client requests here in it's own routine
	go serverss.ServeClients()
	// First time being called fetch feeds
	fetchrss.FetchAndStoreRSSFeeds()
	// Periodically run the fetch and store RSS feeds function every 5 minutes
	updateTimer := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-updateTimer.C:
			fetchrss.FetchAndStoreRSSFeeds()
		}
	}
}
