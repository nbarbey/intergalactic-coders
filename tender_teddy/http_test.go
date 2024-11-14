package economy_of_tests

import (
	"net/http/httptest"
	"testing"
)

func TestHTTP_PublishAd(t *testing.T) {
	app := NewApplication()
	app.Start()
	server := httptest.NewServer(app.http.mux)
	testPublishAd(t, NewClient(server.Client()).WithURL(server.URL))
}

func TestHTTP_SearchAd(t *testing.T) {
	app := NewApplication()
	app.Start()
	server := httptest.NewServer(app.http.mux)
	testSearchAd(t, NewClient(server.Client()).WithURL(server.URL))
}
