package economy_of_tests

import "strings"

type ClassifiedAd struct {
	Title, Body string
	Price       int
}

func (a ClassifiedAd) Matches(word string) bool {
	return strings.Contains(a.Title, word) || strings.Contains(a.Body, word)
}

type Application struct {
	ads []ClassifiedAd
}

func (a *Application) ListAds() ([]ClassifiedAd, error) {
	return a.ads, nil
}

func (a *Application) Publish(ad ClassifiedAd) error {
	a.ads = append(a.ads, ad)
	return nil
}

func (a *Application) Search(word string) (ads []ClassifiedAd, err error) {
	for _, ad := range a.ads {
		if ad.Matches(word) {
			ads = append(ads, ad)
		}
	}
	return ads, nil
}

func NewApplication() *Application {
	return &Application{}
}
