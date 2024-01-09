package types

import "github.com/dave/jennifer/jen"

// packageAliases is a map of package short names to package paths. For use with Jennifer.
var packageAliases = map[string]string{
	"data":     "github.com/ethanmoffat/eolib-go/pkg/eolib/data",
	"net":      "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net",
	"protocol": "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol",
	"pub":      "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/pub",
}

func PackagePath(packageName string) string {
	if v, ok := packageAliases[packageName]; ok {
		return v
	}

	return packageName
}

func AddImports(f *jen.File) {
	for k, v := range packageAliases {
		f.ImportName(v, k)
	}
}
