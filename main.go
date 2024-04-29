package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const Version = "v1.0.0"

func main() {
	fmt.Printf("OpenBinary %s\n\n", Version)
	if len(os.Args) <= 1 {
		fmt.Printf("Please specify file to open with %s\n", os.Args[0])
		os.Exit(1)
	}

	for _, bf := range os.Args[1:] {
		OpenWithPowershell(bf)
	}
}

func OpenWithPowershell(bf string) {
	bfp, err := filepath.Abs(bf)
	if err != nil {
		fmt.Printf("Could not find '%s'; skipping\n", bf)
	}

	fmt.Printf("Opening '%s':\n", bfp)
	cmd := exec.Command("powershell.exe","-Command","start","'" + bfp + "'")
	err = cmd.Start()
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
}