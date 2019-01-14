package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"strings"
)

func printMessage(message string) {
	fmt.Println(message)
}

func help() {
	printMessage("Usage:")
	printMessage("ship (cmd) [options]")
	printMessage("")
	printMessage("ship init              Initializes everything")
	printMessage("ship deploy (template) Deploy template")
}

func execCommand() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func args() {
	argsWithoutProg := os.Args[1:]

	if (len(argsWithoutProg)==0) {
		help()
		os.Exit(100)
	}

	arg1 := os.Args[1]
	switch(arg1) {
	case "init":
		cmd_init();
		break;
	case "deploy":
		cmd_deploy();
		break;
	default:
	}
}

func main() {
	args()
}
