package main

import (
	"fmt"
)

// Original struct
type CityTemperatureOriginal struct {
	City            string
	Temperature     float64
	RegionID        int
	IsCoastal       bool
	PopularityScore int
}

// Optimized struct (arranged for better memory usage)
type CityTemperatureOptimized struct {
	City            string
	IsCoastal       bool
	Temperature     float64
	PopularityScore int
	RegionID        int
}

// Simulated dataset with more cities
var cityTemperatures = map[string]CityTemperatureOriginal{
	"New York":      {"New York", 25.5, 1, false, 2},
	"Los Angeles":   {"Los Angeles", 30.0, 2, false, 1},
	"Chicago":       {"Chicago", 20.0, 3, true, 3},
	"Houston":       {"Houston", 32.5, 4, false, 2},
	"Phoenix":       {"Phoenix", 40.0, 5, false, 1},
	"San Francisco": {"San Francisco", 18.0, 6, true, 4},
	"Miami":         {"Miami", 35.0, 7, true, 5},
	"Seattle":       {"Seattle", 15.0, 8, true, 6},
	"Denver":        {"Denver", 28.0, 9, false, 7},
	"Boston":        {"Boston", 22.0, 10, true, 8},
	"Atlanta":       {"Atlanta", 29.0, 11, true, 9},
	"Philadelphia":  {"Philadelphia", 24.0, 12, false, 10},
}

// Fetch temperature for the original struct with added memory allocation
func fetchTemperatureOriginal(city string) (float64, error) {
	data, ok := cityTemperatures[city]
	if !ok {
		return 0, fmt.Errorf("city not found")
	}
	// Simulate memory allocation
	tempSlice := make([]float64, 1000) // Allocate a larger slice
	for i := 0; i < len(tempSlice); i++ {
		tempSlice[i] = data.Temperature // Use the temperature in the slice
	}
	return data.Temperature, nil
}

// Fetch temperature for the optimized struct with added memory allocation
func fetchTemperatureOptimized(city string) (float64, error) {
	// Create a local slice to simulate an allocation
	tempSlice := make([]float64, 1000) // Allocate a larger slice
	data, ok := cityTemperatures[city]
	if !ok {
		return 0, fmt.Errorf("city not found")
	}
	for i := 0; i < len(tempSlice); i++ {
		tempSlice[i] = data.Temperature
	}
	return data.Temperature, nil
}
