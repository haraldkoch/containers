package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/tplink-plug-exporter:rolling")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{
		Port: "9233",
		Path: "/metrics",
	}, nil)
}
