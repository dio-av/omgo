# Open-Meteo-Go

A simple go client for the [open meteo](https://open-meteo.com) API. 

This fork will provide an implementation for Forecast and Historical API, so the client can get both data of Open Meteo.

## Usage

Simple example of how it'll look like:

```go
package main

import (
	"context"
	"fmt"

	"github.com/hectormalot/omgo"
)

func main() {
	f, _ := omgo.NewClient()

	// Get the current weather for amsterdam
	loc, _ := omgo.NewLocation(52.3738, 4.8910)
	res, _ := f.CurrentWeather(context.Background(), loc, nil)
	fmt.Println("The temperature in Amsterdam is: ", res.Temperature)

	// Get the humidity and cloud cover forecast for berlin, 
	// including the last 2 days and non-metric units
	loc, _ := omgo.NewLocation(52.5235, 13.4115)
	opts := omgo.ForecastOptions{
		TemperatureUnit:   "fahrenheit",
		WindspeedUnit:     "mph",
		PrecipitationUnit: "inch",
		Timezone:          "US/Eastern",
		HourlyMetrics:     []string{"cloudcover, relativehumidity_2m"},
		DailyMetrics:      []string{"temperature_2m_max"},
	}
	
	res, _ := f.Forecast(context.Background(), loc, &opts)
	fmt.Println(res)
	// res.HourlyMetrics["cloudcover"] contains an array of cloud coverage predictions
	// res.HourlyMetrics["relativehumidity_2m"] contains an array of relative humidity predictions
	// res.HourlyTimes contains the timestamps for each prediction
	// res.DailyMetrics["temperature_2m_max"] contains daily maximum values for the temperature_2m metric
	// res.DailyTimes contains the timestamps for all daily predictions
	
	hopts := omgo.HistoricalOptions{
		TemperatureUnit:   "fahrenheit",
		WindspeedUnit:     "mph",
		PrecipitationUnit: "inch",
		Timezone:          "US/Eastern",
		StartDate:         "2023-05-01",
		EndDate:           "2023-06-01",
		HourlyMetrics:     []string{"cloudcover, relativehumidity_2m"},
		DailyMetrics:      []string{"temperature_2m_max", "sunset"},
	}

	hres, _ := f.Historical(context.Background(), loc, &hopts)
	fmt.Println(hres)
}


```
