package main

import (
	"github.com/zserge/webview"
	"net/http"
	"fmt"
	"io/ioutil"
)

func getCity() string {
	getCity, err := http.Get("https://ifconfig.co/city")
	if err != nil {
		fmt.Print(err)
	}
	defer getCity.Body.Close()
	if getCity.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(getCity.Body)
		if err != nil {
			fmt.Print(err)
		}
		return string(bodyBytes)
	}
	return string(getCity.StatusCode)
}

func main()  {
	apiKey := " "
	webview.Open("weatherChecker",
		"http://api.openweathermap.org/data/2.5/weather?q="+getCity()+"&mode=html&appid="+apiKey, 140, 170, true)
}
