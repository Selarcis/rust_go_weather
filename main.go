package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetFile(url string) {

	filename := strings.Split(url, "/")
	out, err := os.Create(filename[len(filename)-1])

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	data, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Body.Close()

	_, err = io.Copy(out, data.Body)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	GetFile("https://www.ncei.noaa.gov/data/global-hourly/access/1901/02907099999.csv")

}
