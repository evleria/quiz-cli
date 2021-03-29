package config

import (
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

var Url = "http://remote.config.com"

func TestRemoteConfig_ReadConfig(t *testing.T) {
	defer gock.Off()

	gock.New(Url).
		Reply(200).
		BodyString(content)

	httpClient := new(http.Client)
	gock.InterceptClient(httpClient)

	config := NewRemoteConfig(httpClient, Url)
	categories, err := config.ReadConfig()

	assert.NilError(t, err)
	assert.Check(t, is.DeepEqual(categories, fixture))
}

func TestRemoteConfig_ReadConfig_InvalidFormat(t *testing.T) {
	defer gock.Off()

	gock.New(Url).
		Reply(200).
		BodyString("ERROR")

	httpClient := new(http.Client)
	gock.InterceptClient(httpClient)

	config := NewRemoteConfig(httpClient, Url)
	_, err := config.ReadConfig()

	assert.ErrorContains(t, err, "yaml: unmarshal errors")
}
