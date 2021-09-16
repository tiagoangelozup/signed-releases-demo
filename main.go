package main

import (
	"fmt"
	"html"
)

func main() {
	emoji := html.UnescapeString("&#128513;")
	fmt.Println("Hi!" + emoji)
}
