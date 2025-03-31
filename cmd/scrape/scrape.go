package scrape

import (
	"encoding/csv"
	species_scrape "github.com/alistairjudson/species-scrape"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "scrape",
		Short: "Scrape for the species",
		Run: func(cmd *cobra.Command, args []string) {
			scraper := &species_scrape.Scraper{URL: "https://www.britishbryologicalsociety.org.uk/learning/species-finder/"}
			species, err := scraper.Scrape()
			if err != nil {
				log.Fatal(err)
			}
			file, err := os.Create("species.csv")
			if err != nil {
				log.Fatal(err)
			}
			headers := []string{
				"ScientificName",
				"Authority",
				"CommonName",
				"Synonyms",
				"Division",
				"GrowthForm",
				"Link",
				"FieldGuide",
				"Distribution",
			}
			csvWriter := csv.NewWriter(file)
			if err := csvWriter.Write(headers); err != nil {
				log.Fatal(err)
			}
			csvWriter.Flush()
			for _, specie := range species {
				if err := csvWriter.Write(specie.Record()); err != nil {
					log.Fatal(err)
				}
				csvWriter.Flush()
			}
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		},
	}
}
