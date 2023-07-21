package xml

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/ethanmoffat/eolib-go/pkg/eolib"
)

type Protocol struct {
	Enums   []ProtocolEnum   `xml:"enum"`
	Structs []ProtocolStruct `xml:"struct"`
	Packets []ProtocolPacket `xml:"packet"`
}

type ProtocolEnum struct {
	Name    string          `xml:"name,attr"`
	Type    string          `xml:"type,attr"`
	Values  []ProtocolValue `xml:"value"`
	Comment string          `xml:"comment"`

	Package     string
	PackagePath string
}

type ProtocolStruct struct {
	Name         string                `xml:"name,attr"`
	Instructions []ProtocolInstruction `xml:",any"`
	Comment      string                `xml:"comment"`

	Package     string
	PackagePath string
}

type ProtocolPacket struct {
	Family       string                `xml:"family,attr"`
	Action       string                `xml:"action,attr"`
	Instructions []ProtocolInstruction `xml:",any"`
	Comment      string                `xml:"comment"`

	Package     string
	PackagePath string
}

type ProtocolValue struct {
	Name    string       `xml:"name,attr"`
	Comment string       `xml:"comment"`
	Value   OrdinalValue `xml:",chardata"`
}

type ProtocolInstruction struct {
	XMLName xml.Name

	// ProtocolField properties
	Name     *string `xml:"name,attr"`
	Type     *string `xml:"type,attr"`
	Length   *string `xml:"length,attr"`
	Padded   *bool   `xml:"padded,attr"`
	Optional *bool   `xml:"optional,attr"`
	Comment  *string `xml:"comment"`
	Content  *string `xml:",chardata"`

	// ProtocolArray properties
	// shared: Name, Type, Length, Optional, Comment
	Delimited         *bool `xml:"delimited,attr"`
	TrailingDelimiter *bool `xml:"trailing-delimiter,attr"`

	// ProtocolLength properties
	// shared: Name, Type, Optional, Comment
	Offset *int `xml:"offset,attr"`

	// ProtocolDummy properties
	// shared: Type, Comment, Content shared

	// ProtocolSwitch properties
	Field *string        `xml:"field,attr"`
	Cases []ProtocolCase `xml:"case"`

	// ProtocolChunked properties
	Chunked []ProtocolInstruction `xml:",any"`

	// ProtocolBreak properties (none)
}

type ProtocolCase struct {
	Value        string                `xml:"value,attr"`
	Default      bool                  `xml:"default,attr"`
	Comment      string                `xml:"comment"`
	Instructions []ProtocolInstruction `xml:",any"`
}

type OrdinalValue int

func validate(instructions []ProtocolInstruction) error {
	for _, inst := range instructions {
		if err := inst.Validate(); err != nil {
			return err
		}

		if err := validate(inst.Chunked); err != nil {
			return err
		}

		if len(inst.Cases) > 0 {
			for _, cs := range inst.Cases {
				if err := validate(cs.Instructions); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (p Protocol) Validate() error {
	for _, st := range p.Structs {
		if err := validate(st.Instructions); err != nil {
			return err
		}
	}

	for _, pkt := range p.Packets {
		if err := validate(pkt.Instructions); err != nil {
			return err
		}
	}

	return nil
}

func (p Protocol) FindType(typeName string) (ps interface{}) {
	if e, ok := p.IsEnum(typeName); ok {
		return e
	}

	if s, ok := p.IsStruct(typeName); ok {
		return s
	}

	return nil
}

func (p Protocol) IsEnum(typeName string) (*ProtocolEnum, bool) {
	for i, e := range p.Enums {
		if e.Name == typeName {
			return &p.Enums[i], true
		}
	}

	return nil, false
}

func (p Protocol) IsStruct(typeName string) (*ProtocolStruct, bool) {
	for i, st := range p.Structs {
		if st.Name == typeName {
			return &p.Structs[i], true
		}
	}

	return nil, false
}

func (p Protocol) IsPacket(typeName string) (*ProtocolPacket, bool) {
	for i, pkt := range p.Packets {
		if pkt.GetTypeName() == typeName {
			return &p.Packets[i], true
		}
	}

	return nil, false
}

func (p ProtocolPacket) GetTypeName() string {
	packageName := string(unicode.ToUpper([]rune(p.Package)[0])) + p.Package[1:]
	return fmt.Sprintf("%s%s%sPacket", p.Family, p.Action, packageName)
}

func (pi ProtocolInstruction) Validate() error {
	all, required, err := pi.expectedFields()

	if err != nil {
		return err
	}

	if len(all) == 0 {
		return nil
	}

	reflectValue := reflect.ValueOf(pi)
	reflectType := reflectValue.Type()

	strEq := func(a string, b string) bool { return a == b }
	for i := 0; i < reflectValue.NumField(); i++ {
		fieldName := reflectType.Field(i)
		fieldValue := reflectValue.Field(i)

		if fieldName.Name == "XMLName" {
			continue
		}

		if eolib.FindIndex(all, fieldName.Name, strEq) == -1 {
			fieldIsNil := fieldValue.IsNil()

			fieldIsEmpty := true
			if !fieldIsNil {
				switch fieldValue.Kind() {
				case reflect.Pointer:
					switch fieldValue.Elem().Kind() {
					case reflect.String:
						fieldIsEmpty = len(strings.TrimSpace(fieldValue.Elem().String())) == 0
					default:
						fieldIsEmpty = fieldValue.Elem().IsZero()
					}
				default:
					fieldIsEmpty = fieldValue.IsZero()
				}
			}

			if !fieldIsNil && !fieldIsEmpty {
				return fmt.Errorf("validation error: instruction of type %s had unexpected field %s with value '%v'", pi.XMLName.Local, fieldName.Name, fieldValue.Elem())
			}
		}

		if eolib.FindIndex(required, fieldName.Name, strEq) >= 0 && fieldValue.IsNil() {
			return fmt.Errorf("valdiation error: instruction of type %s missing required field %s", pi.XMLName.Local, fieldName.Name)
		}
	}

	return nil
}

func (pi ProtocolInstruction) expectedFields() (all []string, required []string, err error) {
	switch pi.XMLName.Local {
	case "field":
		all = []string{"Name", "Type", "Length", "Padded", "Optional", "Comment", "Content"}
		required = []string{"Type"}
	case "array":
		all = []string{"Name", "Type", "Length", "Optional", "Comment", "Delimited", "TrailingDelimiter"}
		required = []string{"Name", "Type"}
	case "length":
		all = []string{"Name", "Type", "Optional", "Comment", "Offset"}
		required = []string{"Name", "Type"}
	case "dummy":
		all = []string{"Type", "Comment", "Content"}
		required = []string{"Type"}
	case "switch":
		all = []string{"Field", "Cases"}
		required = all
	case "chunked":
		all = []string{"Chunked"}
	case "break":
		break
	default:
		err = fmt.Errorf("validation error: invalid xml name '%s'", pi.XMLName.Local)
	}

	return
}

func (o *OrdinalValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	var s string
	if err = d.DecodeElement(&s, &start); err != nil {
		return err
	}

	var val int64
	if val, err = strconv.ParseInt(s, 10, 32); err != nil {
		return err
	}

	*o = OrdinalValue(val)

	return nil
}
