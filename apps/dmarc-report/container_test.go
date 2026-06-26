package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/dmarc-report:rolling")
	helpers.RequireFileExists(t, image, "/usr/bin/dmarcts-report-parser.pl")
}
