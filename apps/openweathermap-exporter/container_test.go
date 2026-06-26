package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/openweathermap-exporter:rolling")
	testhelpers.RequireFileExists(t, image, "/app/openweathermap-exporter", nil)

	// testhelpers.RequireHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{
	// 	Port: "2112",
	// 	Path: "/metrics",
	// }, nil)
}
