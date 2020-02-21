package fetchrss

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// For each of the sources of the rss feeds have a struct for them containing all the XML as a string and whether it's done read
type RSSXML struct {
	Content []byte
}

func GetRSSXML(url string) (*RSSXML, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &RSSXML{Content: body}, nil
}
func StoreFeeds(conn *sql.DB, sourceId int, xmlContent *RSSXML) error {
	transaction, err := conn.Begin()
	if err != nil {
		fmt.Println("Error in creating transaction:> ", err)
	}

	// parse the xmlcontent for this particular rss topic
	feedSourceResults := []CNNRss{}
	// Since NewDecoder needs something that is a reader convert our content []byte into a reader
	newReader := bytes.NewReader(xmlContent.Content)
	decoder := xml.NewDecoder(newReader)
	err = decoder.Decode(&feedSourceResults)
	if err != nil {
		fmt.Println("Error in decoding XML: >", err)
		log.Fatal(err)
	}
	for _, value := range feedSourceResults {
		fmt.Printf("%T", value.Channel.Item, " Is the TYPE of the news Items")
		fmt.Println()
		// For all the news items in the feed items slice loop over and save them if they are newer than last record
		for _, newsItem := range value.Channel.Item {
			title := newsItem.Title
			description := newsItem.Description
			link := newsItem.Link
			guid := newsItem.Guid
			// guid is a struct so it has to be converted to a json string
			guidString, err := json.Marshal(&guid)
			if err != nil {
				log.Fatal(err)
			}
			pubDate := newsItem.PubDate
			// Need to convert this pubDate to someething MySQL likes
			pubDateTime, err := time.Parse(time.RFC1123, pubDate)
			pubDateMysql := pubDateTime.Format("2006-01-02 15:04:05")
			if err != nil {
				log.Fatal(err)
			}
			dateCreated := time.Now().Format("2006-01-02 15:04:05")
			insertS, err := transaction.Prepare("INSERT INTO feeds (source_id, title, description, link, guid, pubDate, dateCreated) VALUES(?,?,?,?,?,?,?)")
			if err != nil {
				fmt.Println("Error in storing feeds :>", err)
				log.Fatal(err)
			}
			defer insertS.Close()
			if _, err := insertS.Exec(sourceId, title, description, link, guidString, pubDateMysql, dateCreated); err != nil {
				transaction.Rollback()
				log.Fatal(err)
			}
		}
	}
	return transaction.Commit()
}
