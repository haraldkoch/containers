package main

import (
	"context"
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/tekton-cli:rolling")
	testhelpers.TestFileExists(t, ctx, image, "/app/bin/tkn", nil)
}
