package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/dmarc-report:rolling")
	testhelpers.RequireFileExists(t, image, "/usr/bin/dmarcts-report-parser.pl")
}
