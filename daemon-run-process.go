package dps

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// daemon daemon running program
var daemon = flag.Bool("d", false, "daemon running program with -d")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *daemon {
		args := os.Args[1:]
		length := len(args)
		for i := 0; i < length; i++ {
			if args[i] == "-d" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("pid", cmd.Process.Pid)
		os.Exit(0)
	}
}
