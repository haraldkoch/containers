package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/prometheus-dnssec-exporter:rolling")
	testhelpers.RequireHTTPEndpoint(t, image, testhelpers.HTTPTestConfig{
		Port: "9204",
		Path: "/metrics",
	}, nil)
}
