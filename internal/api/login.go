package api

import (
	"github.com/leocov-dev/tadpoles-backup/internal/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

func Login(email string, password string) error {
	log.Debug("Login...")
	resp, err := client.ApiClient.PostForm(
		client.LoginEndpoint,
		url.Values{
			"email":    {email},
			"password": {password},
			"service":  {"tadpoles"},
		},
	)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return client.NewRequestError(resp)
	}

	log.Debug("Login successful")
	return nil
}

// Must call admit endpoint before any other requests to get proper auth cookies set
func Admit() error {
	log.Debug("Admit...")
	t := time.Now()
	zone, _ := t.Zone()
	log.Debug("zone: ", zone)
	resp, err := client.ApiClient.PostForm(
		client.AdmitEndpoint,
		url.Values{
			"tz": {zone},
		},
	)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return client.NewRequestError(resp)
	}

	log.Debug("Admit successful")
	return nil
}
