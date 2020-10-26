package main

import (
	"os"

	"github.com/enflow.io/enf2/cmd/enf2d/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
