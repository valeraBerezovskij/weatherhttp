package openweather

import (
	"fmt"
)

type AssetsResponse struct {
	Main    MainData      `json:"main"`
	Weather []weatherData `json:"weather"`
	Name    string        `json:"name"`
}

type MainData struct {
	Temp     float64 `json:"temp"`
	Feels    float64 `json:"feels_like"`
	Humidity float64 `json:"humidity"`
}

func (m *MainData) Info() string {
	return fmt.Sprintf("Temperature: %.2f°C\nFeels like: %.2f°C\nHumidity: %.2f%%\n",
		toCelsius(m.Temp), toCelsius(m.Feels), m.Humidity)

}

type weatherData struct {
	Status      string `json:"main"`
	Description string `json:"description"`
}

func (w *weatherData) Info() string {
	return fmt.Sprintf("Weather: %s (%s)\n", w.Status, w.Description)
}
