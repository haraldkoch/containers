package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/actions-runner:rolling")
	testhelpers.TestFileExists(t, image, "/usr/local/bin/yq", nil)
}
