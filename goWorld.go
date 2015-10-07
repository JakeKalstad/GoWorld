package goWorld

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiVer = "1.0/"
const queryUrl string = "http://api.population.io/" + apiVer
const countryUrl string = queryUrl + "countries"
const mortalityUrl string = queryUrl + "mortality-distribution/"
const populationUrl string = queryUrl + "population/"
const expectancyUrl string = queryUrl + "life-expectancy/"
const rankUrl string = queryUrl + "wp-rank/"

type CountryList struct {
	Countries []string `json:"countries"`
}

type Rank struct {
	Dob     string `json:"dob"`
	Gender  string `json:"sex"`
	Country string `json:"country"`
	Rank    string `json:"rank"`
}

type RankRequest struct {
	Rank
	Age    int32
	OnDate string
	Ago    string
}


func Query(urls []string) []io.ReadCloser {
	ch := make(chan io.ReadCloser, len(urls))
	var responses []io.ReadCloser
	for _, url := range urls {
		fmt.Println("Querying: " + url)
		go func(url string) {
			resp, _ := http.Get(url)
			ch <- resp.Body
		}(url)
	}
	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			fmt.Printf(".")
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func GetRank(r RankRequest) (records []Rank, err error) {
	if r.Dob == "" || r.Gender == "" || r.Country == "" {
		return nil, err
	}
	url := rankUrl + r.Dob + "/" + r.Gender + "/" + r.Country + "/"
	var urls []string
	if r.Age != 0 {
		urls = append(urls, url+"aged/"+string(r.Age))
	}
	if r.OnDate != "" {
		urls = append(urls, url+"on/"+r.OnDate)
	}
	if r.Ago != "" {
		urls = append(urls, url+"ago/"+r.Ago)
	}
	var sets []Rank
	for _, body := range Query(urls) {
		defer body.Close()
		var data Rank
		decoder := json.NewDecoder(body)
		err = decoder.Decode(&data)
		sets = append(sets, data)
	}
	return sets, nil
}

func GetCountries() CountryList {
	var data CountryList
	body := Query([]string{countryUrl})[]
	defer body.Close()
	decoder := json.NewDecoder(body)
	_ = decoder.Decode(&data)
	return data
}
