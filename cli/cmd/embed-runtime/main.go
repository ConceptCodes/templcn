package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	assetPath := filepath.Join(root, "..", "docs-site", "public", "tanstack-runtime.js")
	data, err := os.ReadFile(assetPath)
	if err != nil {
		panic(fmt.Errorf("read TanStack runtime: %w", err))
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	output := filepath.Join(root, "runtime_embed.go")
	content := `package main

import "encoding/base64"

const tanstackRuntimeAssetBase64 = ` + fmt.Sprintf("%q", encoded) + `

var tanstackRuntimeAsset = decodeTanStackRuntime()

func decodeTanStackRuntime() string {
	data, err := base64.StdEncoding.DecodeString(tanstackRuntimeAssetBase64)
	if err != nil {
		panic(err)
	}
	return string(data)
}
`
	if err := os.WriteFile(output, []byte(content), 0644); err != nil {
		panic(fmt.Errorf("write embedded runtime: %w", err))
	}
}
