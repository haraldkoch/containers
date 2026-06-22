package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/ubuntu:rolling")
	testhelpers.TestFileExists(t, image, "/scripts/greeting.sh", nil)
	testhelpers.TestFileExists(t, image, "/scripts/umask.sh", nil)
	testhelpers.TestFileExists(t, image, "/scripts/vpn.sh", nil)
	testhelpers.TestFileExists(t, image, "/usr/local/bin/envsubst", nil)
}
