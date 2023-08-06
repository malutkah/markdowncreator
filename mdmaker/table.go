package mdmaker

import (
	"fmt"
)

func Table(data [][]string) string {
	var result string
	
	fmt.Println(data[0])
	fmt.Println(data[1])
	
	// Create header row
	for _, item := range data[0] {
		result += "| " + item + " "
	}
	
	result += "|" + Break
	
	// Add divider
	for range data[0] {
		result += "|---"
	}
	
	result += "|" + Break
	
	// Create remaining rows
	for _, row := range data[1:] {
		for _, item := range row {
			result += "| " + item + " "
		}
		result += "|" + Break
	}
	
	return result
}
