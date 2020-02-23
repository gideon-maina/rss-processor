//Package search helps in searching teh feeds currently stored in the database and returning matching results
//Relies on the MySQL FULLTEXT index on the title, link and description columns for feeds.
package search

import (
	"database/sql"
	"fmt"
	"log"
)

type Result struct {
	Id              int
	Link            string
	Title           string
	PublicationDate string
	Description     string
	Guid            string
}
type SearchResult struct {
	Query  string
	Result []Result
}

// GetSearchResults takes a DB connection and a search Query and returns a SearchResult, which has the query and a []Result
func GetSearchResults(conn *sql.DB, searchQuery string) (SearchResult, error) {
	results := []Result{}
	sourceRows, err := conn.Query("SELECT id,title,link,guid,description,pubDate FROM feeds WHERE MATCH(title, description, link) AGAINST(?)", searchQuery)
	if err != nil {
		log.Fatal(err)
	}
	for sourceRows.Next() {
		var result Result
		err = sourceRows.Scan(&result.Id, &result.Title, &result.Link, &result.Guid, &result.Description, &result.PublicationDate)
		if err != nil {
			fmt.Println("Failed to search for the query")
			log.Fatal(err)
		}
		results = append(results, result)
	}
	searchResults := SearchResult{Query: searchQuery, Result: results}
	return searchResults, nil
}
