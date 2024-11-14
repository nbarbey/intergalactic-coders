package economy_of_tests

import (
	"encoding/json"
	"io"
	"net/http"
)

type HTTP struct {
	mux *http.ServeMux
}

func NewHTTP(app Ader) *HTTP {
	return &HTTP{mux: withHandlers(app)(http.NewServeMux())}
}

func withHandlers(app Ader) func(mux *http.ServeMux) *http.ServeMux {
	return func(mux *http.ServeMux) *http.ServeMux {
		mux.HandleFunc("GET /list", func(w http.ResponseWriter, r *http.Request) {
			ads, _ := app.ListAds()
			body, _ := json.Marshal(ads)
			_, _ = w.Write(body)
		})
		mux.HandleFunc("POST /ads", func(w http.ResponseWriter, r *http.Request) {
			var ad ClassifiedAd
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &ad)
			_ = app.Publish(ad)
		})
		mux.HandleFunc("GET /ads", func(w http.ResponseWriter, r *http.Request) {
			word := r.URL.Query().Get("word")
			ads, _ := app.Search(word)
			body, _ := json.Marshal(ads)
			_, _ = w.Write(body)
		})
		return mux
	}
}
