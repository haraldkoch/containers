package main

import (
	"context"
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/kubectl:rolling")
	testhelpers.TestFileExists(t, ctx, image, "/usr/local/bin/kubectl", nil)
}
