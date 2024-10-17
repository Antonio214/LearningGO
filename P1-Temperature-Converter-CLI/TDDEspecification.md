# Temperature Converter CLI - Test Specifications

We will use Test-Driven Development (TDD) to build the **Temperature Converter CLI**. Below are the test specifications for the project.

## Functions to Test

1. `CelsiusToFahrenheit(celsius float64) float64`
2. `FahrenheitToCelsius(fahrenheit float64) float64`

---

## Test Cases

### 1. `CelsiusToFahrenheit(celsius float64) float64`

- **Test Case 1**: Convert 0°C to Fahrenheit.
  - **Input**: `0`
  - **Expected Output**: `32`

- **Test Case 2**: Convert 100°C to Fahrenheit.
  - **Input**: `100`
  - **Expected Output**: `212`

- **Test Case 3**: Convert -40°C to Fahrenheit.
  - **Input**: `-40`
  - **Expected Output**: `-40`

- **Test Case 4**: Convert 37°C to Fahrenheit (average human body temperature).
  - **Input**: `37`
  - **Expected Output**: `98.6`

- **Test Case 5**: Convert 25.5°C to Fahrenheit.
  - **Input**: `25.5`
  - **Expected Output**: `77.9`

- **Test Case 6**: Convert absolute zero (-273.15°C) to Fahrenheit.
  - **Input**: `-273.15`
  - **Expected Output**: `-459.67`

---

### 2. `FahrenheitToCelsius(fahrenheit float64) float64`

- **Test Case 1**: Convert 32°F to Celsius.
  - **Input**: `32`
  - **Expected Output**: `0`

- **Test Case 2**: Convert 212°F to Celsius.
  - **Input**: `212`
  - **Expected Output**: `100`

- **Test Case 3**: Convert -40°F to Celsius.
  - **Input**: `-40`
  - **Expected Output**: `-40`

- **Test Case 4**: Convert 98.6°F to Celsius (average human body temperature).
  - **Input**: `98.6`
  - **Expected Output**: `37`

- **Test Case 5**: Convert 77.9°F to Celsius.
  - **Input**: `77.9`
  - **Expected Output**: `25.5`

- **Test Case 6**: Convert absolute zero (-459.67°F) to Celsius.
  - **Input**: `-459.67`
  - **Expected Output**: `-273.15`

---

## Edge Cases and Validation

- **Floating-Point Precision**: Ensure the functions handle floating-point arithmetic correctly within an acceptable margin of error (e.g., ±0.01).
- **Extremely High/Low Values**: Test with extremely high or low temperatures to check for potential overflows or underflows.
- **Invalid Inputs**: While the functions accept `float64`, consider how the program should handle invalid inputs if the functionality is expanded (e.g., NaN, infinity).

---

## Summary Table of Test Cases

| Function                   | Input       | Expected Output |
|----------------------------|-------------|-----------------|
| **CelsiusToFahrenheit**    | `0`         | `32`            |
|                            | `100`       | `212`           |
|                            | `-40`       | `-40`           |
|                            | `37`        | `98.6`          |
|                            | `25.5`      | `77.9`          |
|                            | `-273.15`   | `-459.67`       |
| **FahrenheitToCelsius**    | `32`        | `0`             |
|                            | `212`       | `100`           |
|                            | `-40`       | `-40`           |
|                            | `98.6`      | `37`            |
|                            | `77.9`      | `25.5`          |
|                            | `-459.67`   | `-273.15`       |

---

## Notes

- **Precision Handling**: Due to the nature of floating-point arithmetic, some results may have minor discrepancies. When implementing the tests, use a tolerance level (e.g., `±0.01`) to compare expected and actual results.
- **Test Coverage**: The test cases cover a range of typical and edge-case inputs, including:
  - Standard conversion points (freezing and boiling points of water).
  - Negative temperatures.
  - Decimal temperatures.
  - Absolute zero.
---
