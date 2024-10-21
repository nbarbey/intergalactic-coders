package economy_of_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabase_Get(t *testing.T) {
	d := NewDatabase()

	tx := d.DB.Exec(`insert into classified_ad_rows (id, created_at, title, body, price) values (10 , date(), "New chair", "Almost new, only one leg is missing", 1000)`)
	require.NoError(t, tx.Error)

	ad, err := d.Get(10)
	require.NoError(t, err)
	require.NotNil(t, ad)

	assert.Equal(t, ClassifiedAd{Title: "New chair", Body: "Almost new, only one leg is missing", Price: 1000}, *ad)
}

func TestDatabase_Create(t *testing.T) {
	d := NewDatabase()

	coolChair := ClassifiedAd{Title: "Cool chair", Body: "You can sit on it, it's incredible !", Price: 1000}
	id, err := d.Create(coolChair)
	require.NoError(t, err)

	ad, err := d.Get(id)
	require.NoError(t, err)
	require.NotNil(t, ad)

	assert.Equal(t, coolChair, *ad)
}

func TestDatabase_Update(t *testing.T) {
	d := NewDatabase()

	coolChair := ClassifiedAd{Title: "Cool chair", Body: "You can sit on it, it's incredible !", Price: 1000}
	id, err := d.Create(coolChair)
	require.NoError(t, err)

	coolChair.Price = 900
	err = d.Update(id, coolChair)
	require.NoError(t, err)

	ad, err := d.Get(id)
	require.NoError(t, err)
	require.NotNil(t, ad)

	assert.Equal(t, 900, ad.Price)
}
