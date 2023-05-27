package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "34234",
	// }
	colors["red"] = "blue"
	// way to delete a key in map delete(colors,"red")
	fmt.Println(colors)
}
