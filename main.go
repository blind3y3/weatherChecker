package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"strconv"
	"net/http"
	"io/ioutil"
	"strings"
)

func getCity() string {
	getCity, err := http.Get("https://ifconfig.co/city")
	if err != nil{
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

func getWeather(apiKey string, city string)  {
	url := "https://api.openweathermap.org/data/2.5/weather?q="+strings.Trim(city, "\n")+"&APPID="+apiKey

	r, err := http.Get(url)
	if err != nil{
		fmt.Print(err)
	}
	defer r.Body.Close()
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Print(err)
	}

	getTemp, err := jsonparser.GetUnsafeString(jsonData, "main", "temp") //Парсим температуру, она в Кельвинах
	if err != nil {
		fmt.Print(err)
	}
	tempToFloat, err := strconv.ParseFloat(getTemp, 32)
	if err != nil {
		fmt.Print(err)
	}
	temp := int32(tempToFloat) - 273
	fmt.Print("Температура: ",temp, "°C")

	getPressure, err := jsonparser.GetUnsafeString(jsonData, "main", "pressure")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("\nДавление: ",getPressure + " hPa")

	getWind, err := jsonparser.GetUnsafeString(jsonData, "wind", "speed")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("\nСкорость ветра: ",getWind + "м/с")

	getClouds, err := jsonparser.GetUnsafeString(jsonData, "clouds", "all")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("\nОблачность: ",getClouds + "%")

	getCityName, err := jsonparser.GetUnsafeString(jsonData, "name")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("\nВаш город: ",getCityName)
}

func main() {
	apiKey:=" "
	getWeather(apiKey, getCity())
}
