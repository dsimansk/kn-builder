package main

import (
    "fmt"
    "github.com/dsimansk/kn-builder/pkg"
    "os"
)

func main() {
    if err := pkg.NewKnBuilderCmd().Execute(); err != nil {
        fmt.Printf("ERROR: %v\n", err)
        os.Exit(1)
    }
}
