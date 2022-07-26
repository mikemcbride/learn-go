package main

// maps are like an object in JavaScript, a hash in ruby, or a dict in python.
// maps are key/value pairs. keys must all have the same type and values must all have the same type.

import "fmt"

func main () {
    // ways to initialize empty maps:
    // var colors map[string]string
    // colors := make(map[string]string)

    // initialize with values:
    colors := map[string]string {
        "red": "#ff0000",
        "green": "#00ff00",
        "blue": "#0000ff",
    }

    // add entries:
    // we always have to use the square brace syntax for maps.
    // we cannot use the dot notation like we can in JS or with structs
    colors["white"] = "#ffffff"

    // remove an entry
    delete(colors, "white")

    printMap(colors)
}

// iterating over a map
func printMap(c map[string]string) {
    for key, value := range c {
        fmt.Println("Hex code for", key, "is", value)
    }
}
