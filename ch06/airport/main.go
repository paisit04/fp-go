package main

import (
	"encoding/json"
	"fmt"
	"os"

	"airport/pkg"
)

type Entry struct {
	Airport struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Airport"`

	// time
	Time struct {
		Label     string `json:"Label"`
		Month     uint   `json:"Month"`
		MonthName string `json:"Month Name"`
		Year      uint   `json:"Year"`
	} `json:"Time"`

	// statistics
	Statistics struct {
		NumberOfDelays struct {
			Carrier                int `json:"Carrier"`
			LateAircraft           int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security               int `json:"Security"`
			Weather                int `json:"Weather"`
		} `json:"# of Delays"`
		// Carriers
		Carriers struct {
			Names string `json:"Names"`
			Total int    `json:"Total"`
		} `json:"Carriers"`

		// Flights
		Flights struct {
			Cancelled int `json:"Cancelled"`
			Delayed   int `json:"Delayed"`
			Diverted  int `json:"Diverted"`
			OnTime    int `json:"On Time"`
			Total     int `json:"Total"`
		} `json:"Flights"`

		MinutesDelayed struct {
			Carrier                int `json:"Carrier"`
			LateAircraft           int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security               int `json:"Security"`
			Weather                int `json:"Weather"`
		} `json:"Minutes Delayed"`
	} `json:"Statistics"`
}

func getEntries() []Entry {
	bytes, err := os.ReadFile("./resources/airlines.json")
	if err != nil {
		panic(err)
	}

	var entries []Entry
	err = json.Unmarshal(bytes, &entries)
	if err != nil {
		panic(err)
	}
	return entries
}

func main() {
	entries := getEntries()

	SEA := pkg.Filter(entries, func(entry Entry) bool {
		return entry.Airport.Code == "SEA"
	})
	WeatherDelayHours := pkg.FMap(SEA, func(entry Entry) int {
		return entry.Statistics.MinutesDelayed.Weather / 60
	})
	totalWeatherDealy := pkg.Sum(WeatherDelayHours)
	fmt.Printf("%v\n", totalWeatherDealy)
}
