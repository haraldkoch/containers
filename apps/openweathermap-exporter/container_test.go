package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/openweathermap-exporter:rolling")
	helpers.RequireFileExists(t, image, "/app/openweathermap-exporter")

	// helpers.RequireHTTPEndpoint(t, ctx, image, helpers.HTTPTestConfig{
	// 	Port: "2112",
	// 	Path: "/metrics",
	// }, nil)
}
