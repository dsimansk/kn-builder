package main

import (
	"fmt"
	"github.com/dsimansk/knb/pkg/commands/root"
	"os"
)

func main() {
	if err := root.NewKnBuilderCmd().Execute(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
