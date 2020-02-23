//Pacakge serverss provides a Web Interface to enable clients send search queries and get matching RSS feeds stored in the DB
package serverss

import (
	"encoding/json"
	"github.com/gideon-maina/rss-processor/db"
	"github.com/gideon-maina/rss-processor/search"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	port = ":9000"
)

// Open a server at port 9000 to serve clients requests
func ServeClients() error {
	http.HandleFunc("/search", SearchAndRespond)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("Failed to start, Retrying.")
		log.Fatal("Can't create server", err)
	}
	log.Println("Server running at port ", port, ". Ready to process search requests")
	return nil
}

// Process the request from the customer for searching using the q Query param
func SearchAndRespond(w http.ResponseWriter, r *http.Request) {
	conn := db.Conn()
	defer conn.Close()
	var waitgroup sync.WaitGroup

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	waitgroup.Add(1)
	go func(w http.ResponseWriter, r *http.Request) {
		searchQuery := r.URL.Query().Get("q")
		searchQuery = sanitizeInput(searchQuery)
		log.Println("Searching for ", searchQuery)
		searchResults, err := search.GetSearchResults(conn, searchQuery)
		if err != nil {
			log.Fatal("Can't get results for search query", err)
			return
		}

		jsonResults, err := json.Marshal(searchResults)
		if err != nil {
			log.Fatal(err)
			return
		}
		// Set the Content-Type to json
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonResults)
		waitgroup.Done()
		return
	}(w, r)
	waitgroup.Wait()
}

// Sanitize customer input from the query given, just removes !\ and $ , for now it's a naive approach
func sanitizeInput(q string) string {
	q = strings.Trim(q, "!/\\$")
	return q
}
