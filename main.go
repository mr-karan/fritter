package main

import "fmt"
import "time"
import "os"
import "encoding/json"
import "log"
import "io/ioutil"
import "github.com/fatih/color"

var CONFIGPATH string = "config.json"

type CountryMap struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func toMap(nameCountryJson []CountryMap) map[string]string {
	// Build a config map from JSON:
	nameCountryMap := map[string]string{}
	for _, v := range nameCountryJson {
		nameCountryMap[v.Name] = v.Country
	}
	return nameCountryMap
}

func LoadConfig() []CountryMap {
	// Load JSON Config from a file
	configFile, err := ioutil.ReadFile(CONFIGPATH)
	if err != nil {
		log.Fatal("Error opening config file", err.Error())
		os.Exit(1)
	}

	var nameCountryJson []CountryMap
	json.Unmarshal(configFile, &nameCountryJson)
	return nameCountryJson
}

func getCountry(name string, nameCountryMap map[string]string) string {
	return nameCountryMap[name]
}

func getTimeZone(loc string) time.Time {
	location, _ := time.LoadLocation(loc)
	now := time.Now().In(location)
	return now
}

func checkDayNight(now time.Time) string {
	hour := now.Hour()
	if hour >= 19 || hour <= 7 {
		return "night"
	}
	return "day"
}

func main() {
	nameCountryJson := LoadConfig()
	confmap := toMap(nameCountryJson)
	var friendName string
	if len(os.Args) > 1 {
		friendName = os.Args[1]
	} else {
		os.Exit(1)
	}
	friend := getCountry(friendName, confmap)
	now := getTimeZone(friend)
	formattedNow := now.Format("15:04")
	output := fmt.Sprintf("It is %s ", formattedNow)
	if checkDayNight(now) == "day" {
		emoji := "🌞"
		color.Yellow(output + emoji)
	} else {
		emoji := "🌚"
		color.Cyan(output + emoji)
	}
}
