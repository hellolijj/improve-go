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
	fmt.Println(os.Args)
	local := exec.Command("git", "branch", "-D", os.Args[1])
	remote := exec.Command("git", "push", "origin", "--delete", os.Args[2])

	execAndOutput(local)
	if all {
		execAndOutput(remote)
	}
}

func execAndOutput(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}
	log.Printf("exec command: %s", cmd.String())
	localOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(localOutput))
}
