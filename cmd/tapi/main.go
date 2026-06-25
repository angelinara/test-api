package main

import (
	"fmt"
	"os"

	"github.com/angelinara/test-api/internal/builder"
	"github.com/angelinara/test-api/internal/parser"
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
	case "list":
		listCmd()
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

func listCmd() {
	items, err := parser.ListRequests(".test-api/requests")
	if err != nil {
		fmt.Println("failed to list requests:", err)
		os.Exit(1)
	}
	if len(items) == 0 {
		fmt.Println("no requests found — use the /test-api skill to create one")
		return
	}
	for _, item := range items {
		fmt.Printf("%-20s %-6s %-40s %s\n", item.Name, item.Method, item.URL, item.Description)
	}
}
