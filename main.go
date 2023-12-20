package main
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)
func main() {
	// Check if a file path is provided as an argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
	}
	// Read data from file
	filename := os.Args[1]
	data, err := readDataFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Calculate linear regression line
	slope, intercept := linearRegression(data)
	linearRegressionLine := fmt.Sprintf("y = %.6fx + %.6f", slope, intercept)
	// Calculate Pearson correlation coefficient
	correlationCoefficient := pearsonCorrelation(data)
	// Print results
	fmt.Println("Linear Regression Line:", linearRegressionLine)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlationCoefficient)
}
// Read data from file and return as a slice of float64 values
func readDataFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
// Calculate the linear regression line (slope and intercept)
func linearRegression(data []float64) (float64, float64) {
	n := len(data)
	xSum := 0.0
	ySum := 0.0
	xySum := 0.0
	xSquaredSum := 0.0
	for i, y := range data {
		x := float64(i)
		xSum += x
		ySum += y
		xySum += x * y
		xSquaredSum += x * x
	}
	slope := (float64(n)*xySum - xSum*ySum) / (float64(n)*xSquaredSum - xSum*xSum)
	intercept := (ySum - slope*xSum) / float64(n)
	return slope, intercept
}
// Calculate the Pearson correlation coefficient
func pearsonCorrelation(data []float64) float64 {
	n := len(data)
	xSum := 0.0
	ySum := 0.0
	xySum := 0.0
	xSquaredSum := 0.0
	ySquaredSum := 0.0
	for i, y := range data {
		x := float64(i)
		xSum += x
		ySum += y
		xySum += x * y
		xSquaredSum += x * x
		ySquaredSum += y * y
	}
	numerator := float64(n)*xySum - xSum*ySum
	denominator := math.Sqrt((float64(n)*xSquaredSum-xSum*xSum)*(float64(n)*ySquaredSum-ySum*ySum))
	return numerator / denominator
}