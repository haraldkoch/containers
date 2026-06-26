package main

import (
	"testing"

	helpers "github.com/haraldkoch/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/haraldkoch/postgres-init:rolling")
	helpers.RequireFileExists(t, image, "/usr/local/bin/createdb")
	helpers.RequireFileExists(t, image, "/usr/local/bin/createuser")
	helpers.RequireFileExists(t, image, "/usr/local/bin/psql")
	helpers.RequireFileExists(t, image, "/usr/local/bin/pg_isready")
}
