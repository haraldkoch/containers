package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/calibre-web:rolling")

	t.Run("HTTP endpoint test", func(t *testing.T) {
		helpers.RequireHTTPEndpoint(t, image, testhelpers.HTTPTestConfig{Port: "8083"}, nil)
	})

	t.Run("Check /opt/kepubify/kepubify exists", func(t *testing.T) {
		helpers.RequireFileExists(t, image, "/opt/kepubify/kepubify", nil)
	})

	t.Run("Check /opt/calibre/ebook-convert exists", func(t *testing.T) {
		helpers.RequireFileExists(t, image, "/opt/calibre/ebook-convert", nil)
	})
}
