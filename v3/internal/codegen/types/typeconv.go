package types

import (
	"strings"

	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
)

type ImportInfo struct {
	Package string
	Path    string
}

func ProtocolSpecTypeToGoType(eoType string, currentPackage string, fullSpec xml.Protocol) (goType string, nextImport *ImportInfo) {
	if strings.ContainsRune(eoType, rune(':')) {
		eoType = strings.Split(eoType, ":")[0]
	}

	switch eoType {
	case "byte":
		fallthrough
	case "char":
		fallthrough
	case "short":
		fallthrough
	case "three":
		fallthrough
	case "int":
		return "int", nil
	case "bool":
		return "bool", nil
	case "blob":
		return "[]byte", nil
	case "string":
		fallthrough
	case "encoded_string":
		return "string", nil
	default:
		match := fullSpec.FindType(eoType)
		goType = eoType

		if structMatch, ok := match.(*xml.ProtocolStruct); ok && structMatch.Package != currentPackage {
			nextImport = &ImportInfo{structMatch.Package, structMatch.PackagePath}
		} else if enumMatch, ok := match.(*xml.ProtocolEnum); ok && enumMatch.Package != currentPackage {
			nextImport = &ImportInfo{enumMatch.Package, enumMatch.PackagePath}
		}

		if nextImport != nil {
			if val, ok := packageAliases[nextImport.Package]; ok {
				nextImport.Path = val
			}
		}

		return
	}
}
