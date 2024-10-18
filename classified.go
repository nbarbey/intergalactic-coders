package economy_of_tests

type Application struct {
	*Database
}

func (a *Application) ListAds() ([]ClassifiedAd, error) {
	return a.Database.List()
}

func (a *Application) Publish(ad ClassifiedAd) error {
	_, err := a.Database.Create(ad)
	return err
}

func NewApplication() *Application {
	return &Application{Database: NewDatabase()}
}
