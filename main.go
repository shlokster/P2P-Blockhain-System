package main

import (
	"os"

	"github.com/abdullah1308/P2P-Blockchain-System/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}