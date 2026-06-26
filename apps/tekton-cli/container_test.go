package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/tekton-cli:rolling")
	testhelpers.RequireFileExists(t, image, "/app/bin/tkn", nil)
}
