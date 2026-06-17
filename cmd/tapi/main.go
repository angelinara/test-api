package main

import (
	"fmt"
	"os"
)

func main() {

	 if len(os.Args) < 2 {
        fmt.Println("expected a subcommand")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "init":
        initCmd(os.Args[2:])
    case "new":
        newCmd(os.Args[2:])
    case "run":
        runCmd(os.Args[2:])
	case "scan":
        scanCmd(os.Args[2:])	
    default:
        fmt.Printf("unknown command: %s\n", os.Args[1])
        os.Exit(1)
    }
}

func initCmd(args []string) {

}

func newCmd(args []string) {

}

func runCmd(args []string) {

}

func scanCmd(args []string) {


}