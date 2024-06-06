package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethanmoffat/eolib-go/internal/xml"
)

// SanitizeTypeName sanitizes the type name for serialization. Effectively, this removes the 'Type'
// suffix if present.
func SanitizeTypeName(typeName string) string {
	if strings.HasSuffix(typeName, "Type") {
		return typeName[:len(typeName)-4]
	}
	return typeName
}

// GetInstructionTypeName gets the type name (and byte size, if present) of an instruction.
func GetInstructionTypeName(inst xml.ProtocolInstruction) (typeName string, typeSize string) {
	if inst.Type == nil {
		return
	}

	if strings.ContainsRune(*inst.Type, rune(':')) {
		split := strings.Split(*inst.Type, ":")
		typeName, typeSize = split[0], split[1]
	} else {
		typeName = *inst.Type
	}

	return
}

// CalculateTypeSize gets the size of the named type by recursively evaluating and summing the size
// of the named type's members
func CalculateTypeSize(typeName string, fullSpec xml.Protocol) (res int, err error) {
	structInfo, isStruct := fullSpec.IsStruct(typeName)
	if !isStruct {
		return getPrimitiveTypeSize(typeName, fullSpec)
	}

	flattenedInstList := xml.Flatten((*structInfo).Instructions)
	for _, instruction := range flattenedInstList {
		switch instruction.XMLName.Local {
		case "array":
			fallthrough
		case "field":
			fieldTypeName, fieldTypeSize := GetInstructionTypeName(*instruction)
			if fieldTypeSize != "" {
				fieldTypeName = fieldTypeSize
			}

			if instruction.Length != nil {
				if length, err := strconv.ParseInt(*instruction.Length, 10, 32); err == nil {
					// length is a numeric constant
					if fieldSize, err := CalculateTypeSize(fieldTypeName, fullSpec); err == nil {
						res += fieldSize * int(length)
					} else {
						return 0, err
					}
				} else {
					return 0, fmt.Errorf("instruction length %s must be a fixed size for %s (%s)", *instruction.Length, *instruction.Name, instruction.XMLName.Local)
				}
			} else {
				if nestedSize, err := getPrimitiveTypeSize(fieldTypeName, fullSpec); err != nil {
					return 0, err
				} else {
					res += nestedSize
				}
			}
		case "break":
			res += 1
		}
	}

	return
}

func getPrimitiveTypeSize(fieldTypeName string, fullSpec xml.Protocol) (int, error) {
	switch fieldTypeName {
	case "byte":
		fallthrough
	case "char":
		return 1, nil
	case "short":
		return 2, nil
	case "three":
		return 3, nil
	case "int":
		return 4, nil
	case "bool":
		return 1, nil
	case "blob":
		fallthrough
	case "string":
		fallthrough
	case "encoded_string":
		return 0, fmt.Errorf("cannot get size of %s without fixed length", fieldTypeName)
	default:
		if _, isStruct := fullSpec.IsStruct(fieldTypeName); isStruct {
			return CalculateTypeSize(fieldTypeName, fullSpec)
		} else if e, isEnum := fullSpec.IsEnum(fieldTypeName); isEnum {
			enumTypeName := SanitizeTypeName(e.Type)
			return getPrimitiveTypeSize(enumTypeName, fullSpec)
		} else {
			return 0, fmt.Errorf("cannot get fixed size of unrecognized type %s", fieldTypeName)
		}
	}
}
