package main

import (
	"fmt"
	"os"
)

func printMessage(message string) {
	fmt.Println(message)
}

func help() {
	printMessage("Usage:")
	printMessage("ship (cmd) [options]")
	printMessage("")
	printMessage("ship init              Initializes everything - install docker and prepare compose")
	printMessage("ship deploy (template) Deploy template from list of templates")
	printMessage("ship list [template]   List templates, match template name")
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
	case "list":
		cmd_list();
		break;
	default:
	}
}

func main() {
	args()
	docker_do()
}
