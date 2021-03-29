package config

import (
	"net/http"
)

type RemoteConfig struct {
	httpClient *http.Client
	url        string
}

func NewRemoteConfig(httpClient *http.Client, url string) *RemoteConfig {
	return &RemoteConfig{httpClient, url}
}

func (c *RemoteConfig) ReadConfig() (*Categories, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return readAll(res.Body)
}
