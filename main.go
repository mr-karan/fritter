package main

import "fmt"

import "time"
import "os"
import "encoding/json"
import "log"
import "io/ioutil"

var CONFIG string = "config.json"

type CountryMap struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

// Can use []CountryMap also, but interface {}
// makes `toJson`` a generic function to convert
// any object to a JSON.
func toMap(countryMap []CountryMap) map[string]string {
	// Build a config map:
	confMap := map[string]string{}
	for _, v := range countryMap {
		confMap[v.Name] = v.Country
	}
	return confMap
}

func LoadConfig() []CountryMap {
	configFile, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		log.Fatal("opening config file", err.Error())
		os.Exit(1)
	}

	var country_map []CountryMap
	json.Unmarshal(configFile, &country_map)
	return country_map
}

func checkCountry(name string, confMap map[string]string) string {
	// And then to find values by key:
	/*if v, ok := confMap[name]; ok {
		fmt.Println("found")
	}
	*/
	return confMap[name]
}

func getTimeZone(loc string) time.Time{
	location, _ := time.LoadLocation(loc)	
	now := time.Now().In(location)
	return now
	
}

func main() {
	countryMap := LoadConfig()
	confmap := toMap(countryMap)
	friend:= checkCountry("Karthika", confmap)
	fmt.Println(getTimeZone(friend))
}

