package main

import (
	"fmt"
	"strings"
)

// Struct to store city data
type City struct {
	Name       string
	Temperature float64
	Rainfall    float64
}

// Function to find the city with the highest temperature
func findHighestTemperature(cities []City) City {
	highest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

// Function to find the city with the lowest temperature
func findLowestTemperature(cities []City) City {
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

// Function to calculate average rainfall
func calculateAverageRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter cities by rainfall threshold
func filterCitiesByRainfall(cities []City, threshold float64) []City {
	filteredCities := []City{}
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

// Function to search for a city by name
func searchCityByName(cities []City, name string) (City, bool) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, true
		}
	}
	return City{}, false
}

func main() {
	// Hardcoded data for cities
	cities := []City{
		{"Mumbai", 29.5, 2500},
		{"Delhi", 32.2, 800},
		{"Chennai", 31.0, 1400},
		{"Kolkata", 30.5, 1700},
		{"Bangalore", 28.0, 850},
	}

	// Find city with highest temperature
	highest := findHighestTemperature(cities)
	fmt.Printf("City with highest temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)

	// Find city with lowest temperature
	lowest := findLowestTemperature(cities)
	fmt.Printf("City with lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)

	// Calculate average rainfall
	averageRainfall := calculateAverageRainfall(cities)
	fmt.Printf("Average rainfall: %.2f mm\n", averageRainfall)

	// Filter cities by rainfall threshold
	var threshold float64
	fmt.Print("Enter rainfall threshold (mm): ")
	_, err := fmt.Scan(&threshold)
	if err != nil {
		fmt.Println("Invalid input for threshold.")
		return
	}
	filteredCities := filterCitiesByRainfall(cities, threshold)
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	for _, city := range filteredCities {
		fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
	}

	// Search for a city by name
	var cityName string
	fmt.Print("Enter city name to search: ")
	fmt.Scanln(&cityName)
	result, found := searchCityByName(cities, cityName)
	if found {
		fmt.Printf("City found: %s - Temperature: %.2f°C, Rainfall: %.2f mm\n", result.Name, result.Temperature, result.Rainfall)
	} else {
		fmt.Println("City not found.")
	}
}
