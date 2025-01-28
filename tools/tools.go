//go:build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "istio.io/tools/cmd/protoc-gen-golang-deepcopy"
)
