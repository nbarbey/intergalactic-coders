package unit1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

func TestElastic(t *testing.T) {
	container, err := elasticsearch.Run(context.Background(),
		"docker.elastic.co/elasticsearch/elasticsearch:8.9.0",
		testcontainers.WithHostPortAccess(9200),
		testcontainers.WithEnv(map[string]string{"ELASTIC_PASSWORD": "test_password"}))
	require.True(t, container.IsRunning())
	s := NewStore()
	require.NoError(t, s.Index(ClassifiedAd{Title: "Car"}))

	ads, err := s.Search("Car")
	require.NoError(t, err)

	assert.Equal(t, "Car", ads[0])
}
