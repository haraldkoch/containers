package main

import (
	"context"
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/ubuntu:rolling")
	testhelpers.TestFileExists(t, ctx, image, "/scripts/greeting.sh", nil)
	testhelpers.TestFileExists(t, ctx, image, "/scripts/umask.sh", nil)
	testhelpers.TestFileExists(t, ctx, image, "/scripts/vpn.sh", nil)
	testhelpers.TestFileExists(t, ctx, image, "/usr/local/bin/envsubst", nil)
}
