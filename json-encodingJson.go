//JSON is and encoding of JavaScript values- string, numbers, booleans, arrays, and objects - as Unicode text.

// Converting a Go data structure like movies to JSON is called marshalling:

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//only exported fields are marshaled, which is why we choose capitalized names for all Go field names
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color, omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bissett"}},
}

func main() {
	//Marshal produces a byte slice containing a very long string with no extranous white space.
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//Unmarshaling decodes the JSON data into a slice of structs:

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
