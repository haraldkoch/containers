package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/actions-runner:rolling")
	helpers.RequireFileExists(t, image, "/usr/local/bin/yq")
}
