//go:build tool

package tool

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/magefile/mage"
	_ "github.com/mfridman/tparse"
	_ "golang.org/x/tools/cmd/stringer"
)
