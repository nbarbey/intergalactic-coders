package economy_of_tests

type Ader interface {
	Publisher
	Lister
	Searcher
}

type Publisher interface {
	Publish(ad ClassifiedAd) error
}

type Lister interface {
	ListAds() ([]ClassifiedAd, error)
}
type Searcher interface {
	Search(word string) (ads []ClassifiedAd, err error)
}
