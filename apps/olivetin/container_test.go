package main

import (
	"context"
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/olivetin:rolling")
	testhelpers.TestFileExists(t, ctx, image, "/usr/local/bin/http", nil)
	testhelpers.TestFileExists(t, ctx, image, "/usr/bin/wol", nil)
	testhelpers.TestHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{Port: "1337"}, nil)
}
