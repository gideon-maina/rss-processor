package search

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestSearchReturnsResults(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("Failed with %v in creating mock DB.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "link", "guid", "description", "pubDate"}).
		AddRow(1, "Lewis Hamilton calls for inclusivity after winning Laureus award", "http://rss.cnn.com/~r/rss/edition_motorsport/~3/7VSSq_-1v8M/index.html", "{}", "Description here", "2020-02-18 10:55:41")

	mock.ExpectQuery("^SELECT (.+) FROM feeds*").
		WithArgs("lewis+hamilton").
		WillReturnRows(rows)

	actualResults, err := GetSearchResults(db, "lewis+hamilton")
	expectedResults := SearchResult{Query: "lewis+hamilton", Result: []Result{}}

	if expectedResults.Query != actualResults.Query {
		t.Errorf("Expected %v results got %v", expectedResults, actualResults)
	}
}
