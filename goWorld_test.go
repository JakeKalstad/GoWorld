package main

import "testing"
import "fmt"

func TestGetCountries(t *testing.T) {
	countryList := GetCountries()
	if len(countryList.Countries) != 201 {
		t.Errorf("Got %d, wanted 201", len(countryList.Countries))
	}
	fmt.Println(countryList)
}

func TestGetRank(t *testing.T) {
	var request RankRequest
	request.Dob = "1952-03-11"
	request.Age = 63
	request.Country = "United Kingdom"
	request.Gender = "male"
	request.Ago = "1y"
	request.OnDate = "2001-05-11"
	rankSets, _ := GetRank(request)
	if (len(rankSets) != 3) {
		t.Errorf("Got %d, wanted 3", len(rankSets))
	}
	fmt.Println(rankSets)
}
