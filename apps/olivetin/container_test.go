package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/olivetin:rolling")
	helpers.RequireFileExists(t, image, "/usr/local/bin/http")
	helpers.RequireFileExists(t, image, "/usr/bin/wol")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{Port: "1337"}, nil)
}
