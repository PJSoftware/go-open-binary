package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const Version = "v1.0.1"

func main() {
	fmt.Printf("OpenBinary %s\n\n", Version)
	if len(os.Args) <= 1 {
		fmt.Printf("Please specify file to open with %s\n", os.Args[0])
		os.Exit(1)
	}

	for _, bf := range os.Args[1:] {
		fmt.Printf("Opening '%s'\n", bf)
		OpenWithPowershell(bf)
	}
}

func OpenWithPowershell(bf string) {
	bfp, err := filepath.Abs(bf)
	if err != nil {
		fmt.Printf("> Could not identify '%s'; skipping\n", bf)
		return
	}

	if _, err := os.Stat(bfp); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("> File '%s' not found; skipping\n", bfp)
		return
	}

	fmt.Printf("Opening '%s':\n", bfp)
	cmd := exec.Command("powershell.exe","-Command","start","'" + bfp + "'")
	err = cmd.Start()
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
}