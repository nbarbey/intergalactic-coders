package unit1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabase_Get(t *testing.T) {
	d := NewDatabase()

	tx := d.DB.Exec(`insert into classified_ads (id, created_at, title, body, price) values (10 , date(), "New chair", "Almost new, only one leg is missing", 1000)`)
	require.NoError(t, tx.Error)

	ad, err := d.Get(10)
	require.NoError(t, err)
	require.NotNil(t, ad)

	assert.Equal(t, "New chair", ad.Title)
	assert.Equal(t, "Almost new, only one leg is missing", ad.Body)
	assert.Equal(t, 1000, ad.Price)
}

func TestDatabase_Create(t *testing.T) {
	d := NewDatabase()

	coolChair := ClassifiedAd{Title: "Cool chair", Body: "You can sit on it, it's incredible !", Price: 1000}
	id, err := d.Create(coolChair)
	require.NoError(t, err)

	ad, err := d.Get(id)
	require.NoError(t, err)
	require.NotNil(t, ad)

	assert.Equal(t, "Cool chair", ad.Title)
	assert.Equal(t, "You can sit on it, it's incredible !", ad.Body)
	assert.Equal(t, 1000, ad.Price)
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
