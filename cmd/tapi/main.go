package main

import (
	"fmt"
	"os"

	"github.com/angelinara/test-api/internal/builder"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected a subcommand")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "init":
		initCmd()
	case "new":
		newCmd(os.Args[2:])
	case "run":
		runCmd(os.Args[2:])
	case "scan":
		scanCmd()
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func initCmd() {
	path := ".test-api/requests"
	if _, err := os.Stat(path); err == nil {
		fmt.Println("already initialized")
		return
	}
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("failed to initialize:", err)
		os.Exit(1)
	}
	fmt.Println("initialized test-api project in", path)

}

func newCmd(args []string) {
	_, err := builder.ParseFlags(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runCmd(args []string) {
}

func scanCmd() {
}
