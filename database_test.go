package economy_of_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
