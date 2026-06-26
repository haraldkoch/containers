package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/maddy:rolling")
	testhelpers.RequireFileExists(t, image, "/bin/maddy", nil)
}
