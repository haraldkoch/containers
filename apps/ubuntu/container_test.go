package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/ubuntu:rolling")
	testhelpers.RequireFileExists(t, image, "/scripts/greeting.sh")
	testhelpers.RequireFileExists(t, image, "/scripts/umask.sh")
	testhelpers.RequireFileExists(t, image, "/scripts/vpn.sh")
	testhelpers.RequireFileExists(t, image, "/usr/local/bin/envsubst")
}
