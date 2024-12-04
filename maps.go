package main

import "fmt"

// Struct containing Lat and Long fields
type Vertex struct {
	Lat, Long float64
}

// Create a map with a string key and Vertex value
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literal but keys are needed
	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	// Here Vertex is omitted as this is a shorter way of writing things out
	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	m := make(map[string]int)

	// Deleting an element and m["Answer"] returns the 0 element
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	/*
		ok is a boolean value representing whether key is in the map
		If key is not present, then elem is the zero value for the map's element type.
	*/
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
