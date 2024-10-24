package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type City struct {
	Name        string  `json:"name"`
	Temperature float64 `json:"temperature"`
}

type CityData struct {
	Cities []City `json:"cities"`
}

var log = logrus.New()
var cityData CityData

func main() {
	// Set up structured logging with Logrus
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	// Load city data from a JSON file
	if err := loadCityData("data.json"); err != nil {
		log.WithFields(logrus.Fields{
			"filename": "data.json",
			"error":    err.Error(),
		}).Fatal("Error loading city data")
	}

	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.Writer())) // Log requests to Logrus
	r.Use(gin.Recovery())                     // Recover from panics

	// Endpoint to get the average temperature of all cities
	r.GET("/average-temperature", func(c *gin.Context) {
		avgTemp := calculateAverageTemperature()
		log.WithField("average_temperature", avgTemp).Info("Calculated average temperature")
		c.JSON(http.StatusOK, gin.H{"average_temperature": avgTemp})
	})

	// Endpoint to get the temperature of a specific city
	r.GET("/temperature/:city", func(c *gin.Context) {
		cityName := c.Param("city")
		temp, err := getTemperatureByCity(cityName)
		if err != nil {
			log.WithFields(logrus.Fields{
				"city":  cityName,
				"error": err.Error(),
			}).Error("City not found")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.WithFields(logrus.Fields{
			"city":        cityName,
			"temperature": temp,
		}).Info("Fetched city temperature")
		c.JSON(http.StatusOK, gin.H{"city": cityName, "temperature": temp})
	})

	r.Run(":8081") // Start the server
}

// Load city data from a JSON file
func loadCityData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.WithFields(logrus.Fields{
			"filename": filename,
			"error":    err.Error(),
		}).Error("Failed to open city data file")
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cityData)
	if err != nil {
		log.WithFields(logrus.Fields{
			"filename": filename,
			"error":    err.Error(),
		}).Error("Failed to decode JSON city data")
		return err
	}
	log.WithField("filename", filename).Info("Successfully loaded city data")
	return nil
}

// Calculate the average temperature of all cities
func calculateAverageTemperature() float64 {
	var total float64
	for _, city := range cityData.Cities {
		total += city.Temperature
	}
	return total / float64(len(cityData.Cities))
}

// Get the temperature of a specific city
func getTemperatureByCity(cityName string) (float64, error) {
	for _, city := range cityData.Cities {
		if city.Name == cityName {
			return city.Temperature, nil
		}
	}
	return 0, fmt.Errorf("city not found: %s", cityName) // Return an error if not found
}
