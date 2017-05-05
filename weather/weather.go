package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func query(city string) (weatherData, error) {
	uri := "http://api.openweathermap.org/data/2.5/weather?APPID=" + "123" + "&q=" + city
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=YOUR_API_KEY&q=" + city)

	if err != nil {
		return weatherData{}, err
	}

	return weatherData{}, nil
}
