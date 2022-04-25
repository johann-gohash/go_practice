package main

import (
	"encoding/json"
	"fmt"
)


func unmarshalling(jsonString string) (res map[string]interface{}) {
	fmt.Println("Unmarshalling...")

	// Create the container for the string that will be unmarshalled
	var reading map[string]interface{}

	// The act of unmarshalling.
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	return reading
}

func main() {
	jsonString := `{"name": "battery sensor", 
                    "capacity": 40, 
                    "time": "2019-09-07T15:50-04:00", 
                    "info": {"desc": "A sensor reading"}
                   }`

	// Print it out.
	fmt.Printf("%+v\n", unmarshalling(jsonString))
}
