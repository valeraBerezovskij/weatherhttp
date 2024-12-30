package main

import (
	"fmt"
	"log"
	"time"
	"valera/openweatherClient/openweather"
)

func main() {
	api := "f52b379bdecbe48443eca0ed1a6ba7ae"
	lat, lon := 50.345009, 19.161020
	owClient, err := openweather.NewClient(time.Second * 10)
	if err != nil{
		log.Fatal(err)
	}
	
	mainData, err := owClient.GetAssetMain(api, lat, lon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mainData.Info())

	weatherData, err := owClient.GetAssetWeather(api, lat, lon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(weatherData[0].Info())
}
