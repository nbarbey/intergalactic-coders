package economy_of_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testPublishAd(t *testing.T, app Ader) {
	require.NoError(t, app.Publish(ClassifiedAd{Title: "Blue Jeans"}))

	ads, err := app.ListAds()
	require.NoError(t, err)

	assert.Equal(t, "Blue Jeans", ads[0].Title)
}

func testSearchAd(t *testing.T, app Ader) {
	require.NoError(t, app.Publish(ClassifiedAd{Title: "old spaceship",
		Body: "cannot fly beyond alpha centaury"}))

	ads, err := app.Search("spaceship")
	require.NoError(t, err)

	assert.Equal(t, "old spaceship", ads[0].Title)
}

func TestAppPublishAd(t *testing.T) {
	testPublishAd(t, NewApplication())
}

func TestAppSearchAd(t *testing.T) {
	testSearchAd(t, NewApplication())
}
