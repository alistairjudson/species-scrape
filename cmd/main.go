package main

import (
	"github.com/alistairjudson/species-scrape/cmd/scrape"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	root := &cobra.Command{
		Use: os.Args[0],
	}
	root.AddCommand(scrape.NewCommand())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
