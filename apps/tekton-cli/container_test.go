package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/tekton-cli:rolling")
	helpers.RequireFileExists(t, image, "/app/bin/tkn")
}
