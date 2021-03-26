package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	GOOGLE_API_KEY = "GET YOUR KEY FROM GOOGLE CLOUD PLATFORM"
)

func GetGeoLocation(state string) (float64, float64) {
	url := ("https://maps.googleapis.com/maps/api/geocode/json?address=" + state + "&key=" + GOOGLE_API_KEY)
	response, err := http.Get(url)
	if err != nil {
		panic("Error retreiving response")
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic("Error retreiving response")
	}

	var longitude float64
	var latitude float64
	var values map[string]interface{}

	json.Unmarshal(body, &values)
	for _, v := range values["results"].([]interface{}) {
		for i2, v2 := range v.(map[string]interface{}) {
			if i2 == "geometry" {
				latitude = v2.(map[string]interface{})["location"].(map[string]interface{})["lat"].(float64)
				longitude = v2.(map[string]interface{})["location"].(map[string]interface{})["lng"].(float64)
				break
			}
		}
	}
	return latitude, longitude
}
