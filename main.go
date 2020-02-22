// Main packge of RSS Processor
package main

import (
	"database/sql"
	"fmt"
	"github.com/gideon-maina/rss-processor/fetchrss"
	"github.com/gideon-maina/rss-processor/search"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

func main() {
	var waitgroup sync.WaitGroup
	conn, err := sql.Open("mysql", "root:admin@/rssfeeds")
	if err != nil {
		fmt.Println("Error in db opening :>", err)
		log.Fatal(err)
	}
	defer conn.Close()

	// Get the sources of the feeds from the DB
	sources, err := fetchrss.GetRSSSources(conn)
	if err != nil {
		fmt.Println("Error in getting RSS sources from DB:>", err)
		log.Fatal(err)
	}

	// Range through the sources and fetch their respective RSS xml files
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
			err = fetchrss.StoreFeeds(conn, source.Id, xmlContent)
			if err != nil {
				fmt.Println("Failed to store the RSS feeds")
			}
			waitgroup.Done()
		}()
	}

	waitgroup.Wait()
	fmt.Println("Done Fetching RSS.")
	// Open a server to serve client requests here
	fmt.Println(search.Hello())
}
