package species_scrape

import "strings"

type Species struct {
	ScientificName string
	Authority      string
	CommonName     []string
	Synonyms       []string
	Division       string
	GrowthForm     string
	Link           string
	FieldGuide     string
	Distribution   string
}

func (s Species) Record() []string {
	return []string{
		s.ScientificName,
		s.Authority,
		strings.Join(s.CommonName, ", "),
		strings.Join(s.Synonyms, ", "),
		s.Division,
		s.GrowthForm,
		s.Link,
		s.FieldGuide,
		s.Distribution,
	}
}
