package main

import (
	"net/http"
	"testing"
)

// test if the function can download a file

func TestGetFile(t *testing.T) {
	url := "https://www.ncei.noaa.gov/data/global-hourly/access/1901/02907099999.csv"
	_, err := http.Get(url)
	if err != nil {
		t.Errorf("Main function for downloading file incountered an error: %v want %v", err, url)
	}
}
