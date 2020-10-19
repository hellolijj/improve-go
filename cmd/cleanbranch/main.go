package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var all bool

func main() {
	flag.BoolVar(&all, "a", false, "delete")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("no params")
		return
	}

	if all {
		execAndOutput(exec.Command("git", "branch", "-D", os.Args[2]))
		execAndOutput(exec.Command("git", "push", "origin", "--delete", os.Args[2]))
	} else {
		execAndOutput(exec.Command("git", "branch", "-D", os.Args[1]))
	}
}

func execAndOutput(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}
	log.Printf("exec command: %s", cmd.String())
	localOutput, _ := cmd.Output()
	fmt.Println(string(localOutput))
}
