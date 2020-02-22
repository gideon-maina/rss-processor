package main

import (
	"fmt"
	"github.com/gideon-maina/rss-processor/fetchrss"
	"github.com/gideon-maina/rss-processor/serverss"
	"log"
	"sync"
)

func main() {
	var waitgroup sync.WaitGroup

	// Get the sources of the feeds from the DB
	sources, err := fetchrss.GetRSSSources()
	if err != nil {
		fmt.Println("Error in getting RSS sources from DB:>", err)
		log.Fatal(err)
	}

	// Range through the sources and fetch their respective RSS xml files concurrently
	for _, sourceVal := range sources {
		source := sourceVal // To enable concurrency force variable evaluation each time in loop
		log.Println("--------------------------- Getting RSS Data for ", source.Url, "-------------------------------")
		go func() {
			waitgroup.Add(1)
			xmlContent, err := fetchrss.GetRSSXML(source.Url)
			if err != nil {
				fmt.Println("Error ranging sources :>", err)
			}
			// Store the rss feeds in the DB
			err = fetchrss.StoreFeeds(source.Id, xmlContent)
			if err != nil {
				fmt.Println("Failed to store the RSS feeds")
			}
			waitgroup.Done()
		}()
	}

	waitgroup.Wait()
	log.Println("Done Fetching RSS.")
	// Open a server to serve client requests here
	serverss.ServeClients()
}
