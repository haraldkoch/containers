package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/openweathermap-exporter:rolling")
	testhelpers.TestFileExists(t, image, "/app/openweathermap-exporter", nil)

	// testhelpers.TestHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{
	// 	Port: "2112",
	// 	Path: "/metrics",
	// }, nil)
}
