package main

import "fmt"

func main() {
	// map of maps
	nested := map[string]map[int]string{
		"mobile": {
			1: "samsung",
			2: "apple",
			3: "xiaomi",
		},
		"tv": {
			1: "panasonic",
			2: "tcl",
			3: "lg",
		},
		"car": {
			1: "tesla",
			2: "nio",
			3: "luga",
		},
	}
	fmt.Println(nested["mobile"][1])
	fmt.Println(nested["tv"][2])
	fmt.Println(nested["car"][3])
}
