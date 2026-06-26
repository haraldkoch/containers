package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/smartctl_exporter:rolling")
	testhelpers.RequireHTTPEndpoint(t, image, testhelpers.HTTPTestConfig{
		Port: "9633",
		Path: "/metrics",
	}, nil)
}
