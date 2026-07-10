//go:build !prod
// +build !prod

package main

import (
	"io/fs"
	"trainspotter-frontend/internal/embedded"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
