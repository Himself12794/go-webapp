package main

import (
	"encoding/json"
	"fmt"
	//"io"
	"log"
	"net/http"
	//"os"
)

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"` // In Kelvins
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"` // In Kelvins
		TempMax  float64 `json:"temp_max"` // |
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt   int    `json:"dt"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func (wr *WeatherResponse) TempToFahrenheit() {
	wr.Main.Temp = KelvinsToFahrenheit(wr.Main.Temp)
	wr.Main.TempMin = KelvinsToFahrenheit(wr.Main.TempMin)
	wr.Main.TempMax = KelvinsToFahrenheit(wr.Main.TempMax)
}

func KelvinsToFahrenheit(t float64) float64 {
	return t*1.8 - 459.67
}

func main() {

	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=4466033&APPID=85383b7344bff8b9af7f717eae540519")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		var resp WeatherResponse
		dec := json.NewDecoder(response.Body)
		dec.Decode(&resp)
		resp.TempToFahrenheit()
		fmt.Printf("%+v", resp)
	}
	fmt.Printf("\n")
}
