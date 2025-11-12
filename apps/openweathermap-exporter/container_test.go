package main

import (
	"context"
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/openweathermap-exporter:rolling")
	testhelpers.TestFileExists(t, ctx, image, "/app/openweathermap-exporter", nil)

	// testhelpers.TestHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{
	// 	Port: "2112",
	// 	Path: "/metrics",
	// }, nil)
}
