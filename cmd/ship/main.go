package main

import (
	"fmt"
)

func printMessage(message string) {
	fmt.Println(message)
}

func help() {
	printMessage("Usage:")
	printMessage("ship type [options]")
}

func main() {
	help()
}
