package economy_of_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListClassifiedAd_no_ads_yet(t *testing.T) {
	app := NewApplication()

	ads, err := app.ListAds()
	require.NoError(t, err)

	assert.Empty(t, ads)
}

func TestPublishAd(t *testing.T) {
	app := NewApplication()

	require.NoError(t, app.Publish(ClassifiedAd{Title: "Blue Jeans"}))

	ads, err := app.ListAds()
	require.NoError(t, err)

	assert.Len(t, ads, 1)
	assert.Equal(t, "Blue Jeans", ads[0].Title)
}
