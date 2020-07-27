package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type LogsRequest struct {
	IP        string    `json:"ip"`
	URL       string    `json:"url"`
	Timestamp time.Time `json:"timestamp"`
}

type API struct {
	Cache   Сache
	Metrics Metrics
}

func New(cache Сache, metrics Metrics) *API {
	return &API{
		Cache:   cache,
		Metrics: metrics,
	}
}

func (api *API) HandleLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errMsg := "Method Not Allowed"

		log.Print("Method Not Allowed")

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(errMsg))

		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Cannot read request body: %v, err: %v", body, err)

		log.Print(errMsg)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))

		return
	}
	defer r.Body.Close()

	var lr LogsRequest
	err = json.Unmarshal(body, &lr)
	if err != nil {
		errMsg := fmt.Sprintf("Cannot unmarshal request body: %v, err: %v", body, err)

		log.Print(errMsg)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))

		return
	}

	err = api.Cache.Set(lr.IP)
	if err != nil {
		if _, ok := err.(ErrValExists); ok {
			errMsg := "Ip adress already exists"

			log.Print(errMsg)

			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(errMsg))

			return
		}

		errMsg := "Error with cache, cannot write"

		log.Print(errMsg)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))

		return
	}

	api.Metrics.Inc()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))

	return
}

func (api *API) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errMsg := "Method Not Allowed"

		log.Print("Method Not Allowed")

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(errMsg))

		return
	}

	promhttp.Handler().ServeHTTP(w, r)
	return
}
