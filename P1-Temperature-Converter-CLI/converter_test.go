package converter

import (
	"math"
	"testing"
)

// Helper function to compare floating-point numbers with a tolerance
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"Freezing Point", 0, 32},
		{"Boiling Point", 100, 212},
		{"Negative Value", -40, -40},
		{"Body Temperature", 37, 98.6},
		{"Decimal Value", 25.5, 77.9},
		{"Absolute Zero", -273.15, -459.67},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToFahrenheit(tt.celsius)
			if !almostEqual(result, tt.expected, 0.01) {
				t.Errorf("CelsiusToFahrenheit(%v) = %v; want %v", tt.celsius, result, tt.expected)
			}
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name       string
		fahrenheit float64
		expected   float64
	}{
		{"Freezing Point", 32, 0},
		{"Boiling Point", 212, 100},
		{"Negative Value", -40, -40},
		{"Body Temperature", 98.6, 37},
		{"Decimal Value", 77.9, 25.5},
		{"Absolute Zero", -459.67, -273.15},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			result := FahrenheitToCelsius(tt.fahrenheit)
			if !almostEqual(result, tt.expected, 0.01) {
				t.Errorf("FahrenheitToCelsius(%v) = %v; want %v", tt.fahrenheit, result, tt.expected)
			}
		})
	}
}
