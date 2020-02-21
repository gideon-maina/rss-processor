// Main packge of RSS Processor
package main

import (
	"database/sql"
	"fmt"
	"github.com/gideon-maina/rss-processor/fetchrss"
	"github.com/gideon-maina/rss-processor/search"
	_ "github.com/go-sql-driver/mysql"
)

type Source struct {
	id            int
	publisher     string
	url           string
	topic         string
	description   string
	lastBuildDate string
	dateCreated   string
	dateModified  string
}
type Feed struct {
	id           int
	sourceId     int
	title        string
	description  string
	link         string
	guid         string
	pubDate      string
	dateCreated  string
	dateModified string
}

func main() {
	fmt.Println(search.Hello())
	conn, err := sql.Open("mysql", "root:admin@/rssfeeds")
	if err != nil {
		fmt.Println("Error in db opening :>", err)
	}
	defer conn.Close()
	// Get the sources of the feeds from the DB
	sources := []Source{}
	sourceRows, err := conn.Query("SELECT id,publisher,url,topic,description,lastBuildDate,dateModified,dateCreated FROM sources;")
	for sourceRows.Next() {
		var source Source
		err = sourceRows.Scan(&source.id, &source.publisher, &source.url, &source.topic, &source.description, &source.lastBuildDate, &source.dateModified, &source.dateCreated)
		sources = append(sources, source)
	}
	// Range through the sources and fetch their respective RSS xml files
	for _, source := range sources {
		fmt.Println("--------------------------- Data for ", source.url, "--------------------------")
		xmlContent, err := fetchrss.GetRSSXML(source.url)
		if err != nil {
			fmt.Println("Error ranging sources :>", err)
		}
		// Store the rss feeds in the DB
		err = fetchrss.StoreFeeds(conn, source.id, xmlContent)
		if err != nil {
			fmt.Println("Failed to store the RSS feeds")
		}
	}
}
