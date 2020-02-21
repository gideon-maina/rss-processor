// Main packge of RSS Processor
package main

import (
	"database/sql"
	"fmt"
	"github.com/gideon-maina/rss-processor/fetchrss"
	"github.com/gideon-maina/rss-processor/search"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "root:admin@/rssfeeds")
	if err != nil {
		fmt.Println("Error in db opening :>", err)
	}
	defer conn.Close()
	// Get the sources of the feeds from the DB
	sources := []fetchrss.Source{}
	sourceRows, err := conn.Query("SELECT id,publisher,url,topic,description,lastBuildDate,dateModified,dateCreated FROM sources;")
	for sourceRows.Next() {
		var source fetchrss.Source
		err = sourceRows.Scan(&source.Id, &source.Publisher, &source.Url, &source.Topic, &source.Description, &source.LastBuildDate, &source.DateModified, &source.DateCreated)
		sources = append(sources, source)
	}
	// Range through the sources and fetch their respective RSS xml files
	for _, source := range sources {
		fmt.Println("--------------------------- Getting RSS Data for ", source.Url, "-------------------------------")
		xmlContent, err := fetchrss.GetRSSXML(source.Url)
		if err != nil {
			fmt.Println("Error ranging sources :>", err)
		}
		// Store the rss feeds in the DB
		err = fetchrss.StoreFeeds(conn, source.Id, xmlContent)
		if err != nil {
			fmt.Println("Failed to store the RSS feeds")
		}
	}
	fmt.Println("Done Fetching RSS.")
	// Open a server to serve client requests here
	fmt.Println(search.Hello())
}
