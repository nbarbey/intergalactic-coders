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
	http *HTTP
	*UseCases
}

type UseCases struct {
	ads []ClassifiedAd
}

func (u *UseCases) ListAds() ([]ClassifiedAd, error) {
	return u.ads, nil
}

func (u *UseCases) Publish(ad ClassifiedAd) error {
	u.ads = append(u.ads, ad)
	return nil
}

func (u *UseCases) Search(word string) (ads []ClassifiedAd, err error) {
	for _, ad := range u.ads {
		if ad.Matches(word) {
			ads = append(ads, ad)
		}
	}
	return ads, nil
}

func (a *Application) Start() {

}

func NewApplication() *Application {
	u := &UseCases{}
	return &Application{
		UseCases: u,
		http:     NewHTTP(u),
	}
}
