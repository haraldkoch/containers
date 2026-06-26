package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/smartctl_exporter:rolling")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{
		Port: "9633",
		Path: "/metrics",
	}, nil)
}
