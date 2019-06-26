package api

import "testing"

func TestBookToJSON(t *testing.T) {
	testBook := Book{Title: "Cloud Native Go", Author: "Vinny Sabatini", ISBN: "8675309"}
	got := testBook.ToJSON()
	expected := `{"title":"Cloud Native Go","author":"Vinny Sabatini","isbn":"8675309"}`

	if string(got) != expected {
		t.Errorf("\nExpected: %s \nGot: %s", expected, got)
	}
}

func TestBookFromJSON(t *testing.T) {
	bookJSON := []byte(`{"title":"Cloud Native Go","author":"Vinny Sabatini","isbn":"8675309"}`)
	got := FromJSON(bookJSON)
	expected := Book{Title: "Cloud Native Go", Author: "Vinny Sabatini", ISBN: "8675309"}

	if got != expected {
		t.Errorf("\nExpected: %s \nGot: %s", expected, got)
	}
}
