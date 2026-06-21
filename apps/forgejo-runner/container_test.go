package main

import (
	"testing"

	"github.com/haraldkoch/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/haraldkoch/forgejo-runner:rolling")

	t.Run("Check cosign exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/usr/local/bin/cosign", nil)
	})

	t.Run("Check flux exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/usr/local/bin/flux", nil)
	})

	t.Run("Check flux-local exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/root/.local/bin/flux-local", nil)
	})

	t.Run("Check flate exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/usr/local/bin/flate", nil)
	})

	t.Run("Check helm exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/usr/local/bin/helm", nil)
	})

	t.Run("Check kustomize exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/usr/local/bin/kustomize", nil)
	})

	t.Run("Check rsync exists", func(t *testing.T) {
		testhelpers.TestFileExists(t, image, "/usr/bin/rsync", nil)
	})
}
