package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/kubectl:rolling")
	testhelpers.RequireFileExists(t, image, "/usr/local/bin/kubectl")
}
