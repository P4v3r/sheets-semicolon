package main

import (
	"os"

	"github.com/P4v3r/sheets/internal/sheets"
)

func main() {
	os.Exit(sheets.Main(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
