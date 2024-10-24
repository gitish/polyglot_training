package main

import (
	"testing"
)

// Benchmark tests with a loop to increase allocations
func BenchmarkMemoryUsageOriginal(b *testing.B) {
	city := "New York"
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ { // Increase number of allocations
			_, err := fetchTemperatureOriginal(city)
			if err != nil {
				b.Errorf("Failed to fetch temperature for city %s: %v", city, err)
			}
		}
	}
}

func BenchmarkMemoryUsageOptimized(b *testing.B) {
	city := "New York"
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ { // Increase number of allocations
			_, err := fetchTemperatureOptimized(city)
			if err != nil {
				b.Errorf("Failed to fetch temperature for city %s: %v", city, err)
			}
		}
	}
}
