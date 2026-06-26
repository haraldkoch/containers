package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/forgejo-runner:rolling")

	t.Run("Check cosign exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/usr/local/bin/cosign")
	})

	t.Run("Check flux exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/usr/local/bin/flux")
	})

	t.Run("Check flux-local exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/root/.local/bin/flux-local")
	})

	t.Run("Check flate exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/usr/local/bin/flate")
	})

	t.Run("Check helm exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/usr/local/bin/helm")
	})

	t.Run("Check kustomize exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/usr/local/bin/kustomize")
	})

	t.Run("Check rsync exists", func(t *testing.T) {
		testhelpers.RequireFileExists(t, image, "/usr/bin/rsync")
	})
}
