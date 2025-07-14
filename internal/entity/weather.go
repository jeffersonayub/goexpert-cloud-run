package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Temp_C float64 `json:"temp_c"`
		Temp_F float64 `json:"temp_f"`
	} `json:"current"`
}

type Response struct {
	Temp_C float64 `json:"temp_C"`
	Temp_F float64 `json:"temp_F"`
	Temp_K float64 `json:"temp_K"`
}

const weather_key = "WEATHER_API_KEY" // Replace with your actual API key

func GetWeather(location string) (Response, error) {
	var weather Weather
	err := fetchWeatherData(location, &weather)
	if err != nil {
		return Response{}, err
	}
	return weather.ToResponse(), nil
}

func (w Weather) ToResponse() Response {
	return Response{
		Temp_C: w.Current.Temp_C,
		Temp_F: w.Current.Temp_F,
		Temp_K: w.Current.Temp_C + 273.15,
	}
}
func fetchWeatherData(location string, weather *Weather) error {
	response, err := http.Get("https://api.weatherapi.com/v1/current.json?key=" + weather_key + "&q=" + url.QueryEscape(location))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch weather data: %s", response.Status)
	}

	return json.NewDecoder(response.Body).Decode(weather)
}
