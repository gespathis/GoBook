package main

import (
	"GoBook/Temperature/converters"
	"fmt"
)

func drawTempTable(start float64, end float64, captionFrom string, captionTo string, tempConverter func(temp float64) float64) {

	fmt.Printf("====================\n")
	fmt.Printf("| %-6v |  %-6v |\n", captionFrom, captionTo)
	fmt.Printf("====================\n")
	for index := start; index <= end; index += 5 {
		fmt.Printf("| %6.2f |  %6.2f |\n", index, tempConverter(index))
	}
	fmt.Printf("====================\n")

}

func main() {
	//start := converters.Celcious(-40.00)
	//end := converters.Celcious(100.00)

	var f = func(temp float64) float64 {
		return float64(converters.Celcious(temp).Fahrenheit())
	}

	drawTempTable(-40.00, 100, "°C", "°F", f)

	var f3 = func(temp float64) float64 {
		return float64(converters.Celcious(temp).Kelvin())
	}

	drawTempTable(-40.00, 212.00, "°C", "°K", f3)

	var f2 = func(temp float64) float64 {
		return float64(converters.Fahrenheit(temp).Celcious())
	}

	drawTempTable(-40.00, 212.00, "°F", "°C", f2)

}
