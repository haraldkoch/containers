package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/olivetin:rolling")
	testhelpers.RequireFileExists(t, image, "/usr/local/bin/http", nil)
	testhelpers.RequireFileExists(t, image, "/usr/bin/wol", nil)
	testhelpers.RequireHTTPEndpoint(t, image, testhelpers.HTTPTestConfig{Port: "1337"}, nil)
}
