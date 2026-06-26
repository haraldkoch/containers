package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/ubuntu:rolling")
	helpers.RequireFileExists(t, image, "/scripts/greeting.sh")
	helpers.RequireFileExists(t, image, "/scripts/umask.sh")
	helpers.RequireFileExists(t, image, "/scripts/vpn.sh")
	helpers.RequireFileExists(t, image, "/usr/local/bin/envsubst")
}
