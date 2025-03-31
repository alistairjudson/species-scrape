package species_scrape

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type Scraper struct {
	URL string
}

func (s *Scraper) Scrape() (specieses []Species, err error) {
	res, err := http.Get(s.URL)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	path := `#classification-filter > dl > dd > p > a`
	for _, selection := range doc.Find(path).EachIter() {
		link, ok := selection.Attr("href")
		if !ok {
			return
		}
		page, pageErr := http.Get(link)
		if pageErr != nil {
			return nil, pageErr
		}
		defer func() {
			if closeErr := page.Body.Close(); closeErr != nil {
				err = errors.Join(err, closeErr)
			}
		}()
		pageDoc, docErr := goquery.NewDocumentFromReader(page.Body)
		if docErr != nil {
			return nil, docErr
		}
		selector := `body > div.main-container > aside > div > div > div > div`
		species := Species{
			Link: link,
		}
		pageDoc.Find(selector).Each(func(i int, s *goquery.Selection) {
			title := strings.TrimSpace(s.Find("h6").Text())
			value := strings.TrimSpace(s.Find("div").Text())
			switch title {
			case "Scientific name":
				species.ScientificName = value
			case "Authority":
				species.Authority = value
			case "Common name":
				species.CommonName = strings.Split(value, "\n")
			case "Synonyms":
				species.Synonyms = strings.Split(value, "\n")
			case "Division":
				species.Division = value
			case "Growth form":
				species.GrowthForm = value
			}
		})
		fieldGuideSelector := `body > div.main-container > main > div:nth-child(1) > a`
		pageDoc.Find(fieldGuideSelector).Each(func(i int, s *goquery.Selection) {
			fieldGuideLink, ok := s.Attr("href")
			if !ok {
				return
			}
			species.FieldGuide = fieldGuideLink
		})

		distributionSelector := `body > div.main-container > main > div:nth-child(2) > a`
		pageDoc.Find(distributionSelector).Each(func(i int, s *goquery.Selection) {
			distributionLink, ok := s.Attr("href")
			if !ok {
				return
			}
			species.Distribution = distributionLink
		})
		specieses = append(specieses, species)
	}
	return specieses, nil
}
