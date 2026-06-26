package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/kubectl:rolling")
	helpers.RequireFileExists(t, image, "/usr/local/bin/kubectl")
}
