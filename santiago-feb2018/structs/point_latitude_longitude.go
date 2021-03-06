package main

import (
	"encoding/json"
	"fmt"

	"github.com/ifreddyrondon/go-talks/santiago-feb2018/structs/point"
)

func main() {
	// START OMIT
	in := []byte(`{"latitude": 1.1, "longitude": 2.2}`)
	// END OMIT
	var p point.Point
	if err := json.Unmarshal(in, &p); err != nil {
		panic(err)
	}
	// pretty print
	response, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(response))
}
