package unit1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Store struct {
	client *elasticsearch.Client
}

func NewStore() *Store {
	client, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9220"},
		Username:  "elastic",
		Password:  "test_password",
	})
	_, _ = client.Indices.Create("classified_ads_index")
	return &Store{client: client}
}

func (s *Store) Index(ad ClassifiedAd) error {
	data, _ := json.Marshal(ad)
	_, err := s.client.Index("classified_ads_index", bytes.NewReader(data))
	return err
}

func (s *Store) Search(word string) (ads []ClassifiedAd, err error) {
	query := fmt.Sprintf(`{ "query": { "query_string": {"query": %s} } }`, word)
	response, err := s.client.Search(
		s.client.Search.WithIndex("classified_ads_index"),
		s.client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &ads)
	return ads, err
}
