package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type City struct {
	Name        string  `json:"name"`
	Temperature float64 `json:"temperature"`
}

type CityData struct {
	Cities []City `json:"cities"`
}

var cityData CityData

func main() {
	r := gin.Default()

	// Poor logging practice: unstructured logging
	r.Use(func(c *gin.Context) {
		log := fmt.Sprintf("Request: %s %s", c.Request.Method, c.Request.URL)
		fmt.Println(log) // Using fmt.Println for logging, not structured
		c.Next()
	})

	// Panic for error handling
	r.GET("/average-temperature", func(c *gin.Context) {
		if err := loadCityData("data.json"); err != nil {
			panic(err) // Using panic instead of proper error handling
		}

		avgTemp := calculateAverageTemperature()
		c.JSON(http.StatusOK, gin.H{"average_temperature": avgTemp})
	})

	// Retrieve temperature without error handling
	r.GET("/temperature/:city", func(c *gin.Context) {
		cityName := c.Param("city")
		temp := getTemperatureByCity(cityName) // No error handling
		c.JSON(http.StatusOK, gin.H{"city": cityName, "temperature": temp})
	})

	r.Run(":8080")
}

func loadCityData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&cityData)
}

func calculateAverageTemperature() float64 {
	var total float64
	for _, city := range cityData.Cities {
		total += city.Temperature
	}
	return total / float64(len(cityData.Cities))
}

func getTemperatureByCity(cityName string) float64 {
	for _, city := range cityData.Cities {
		if city.Name == cityName {
			return city.Temperature
		}
	}
	return 0 // Returns 0 if the city is not found
}
