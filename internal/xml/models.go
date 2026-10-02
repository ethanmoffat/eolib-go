package xml

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/ethanmoffat/eolib-go/v3/utils"
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
	// XMLName is the XML element name of this instruction.
	XMLName xml.Name
	// IsChunked is True if this instruction appears within a "chunked" section
	IsChunked bool
	// ReferencedBy is the name of the instruction that references this instruction. This is only set for length instructions which are referenced by another instruction.
	ReferencedBy *string

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
	IsChunked bool

	Value        string                `xml:"value,attr"`
	Default      bool                  `xml:"default,attr"`
	Comment      string                `xml:"comment"`
	Instructions []ProtocolInstruction `xml:",any"`
}

type OrdinalValue int

func Flatten(instructions []ProtocolInstruction) []*ProtocolInstruction {
	var flattenedInstList []*ProtocolInstruction
	for ndx, instruction := range instructions {
		if instruction.XMLName.Local == "chunked" {
			flattenedInstList = append(flattenedInstList, Flatten(instruction.Chunked)...)
		} else {
			// note: when flattening, switches are *not* flattened as they may have cases with instructions with name collisions
			// see: CharacterReplyServerPacket
			flattenedInstList = append(flattenedInstList, &instructions[ndx])
		}
	}
	return flattenedInstList
}

func getLengthInstructions(instructions []ProtocolInstruction) (lengthInstructions []*ProtocolInstruction) {
	flattened := Flatten(instructions)
	for i, inst := range flattened {
		if inst.XMLName.Local == "length" {
			lengthInstructions = append(lengthInstructions, flattened[i])
		}
	}
	return
}

func findLengthInstructionByName(lengthName string, lengthInstructions []*ProtocolInstruction) *ProtocolInstruction {
	for i, inst := range lengthInstructions {
		if inst.Name != nil && *inst.Name == lengthName {
			return lengthInstructions[i]
		}
	}
	return nil
}

func validate(instructions []ProtocolInstruction, isChunked bool, lengthInstructions []*ProtocolInstruction) error {
	localLengths := append(lengthInstructions, getLengthInstructions(instructions)...)

	for i, inst := range instructions {
		if isChunked {
			instructions[i].IsChunked = true
		}

		if inst.Length != nil {
			if lengthInstruction := findLengthInstructionByName(*inst.Length, localLengths); lengthInstruction != nil {
				if lengthInstruction.ReferencedBy != nil {
					return fmt.Errorf("validation error: length field %s must not be referenced by multiple fields (%s, %s)", *inst.Length, *lengthInstruction.ReferencedBy, *inst.Name)
				}
				lengthInstruction.ReferencedBy = new(string)
				*lengthInstruction.ReferencedBy = *inst.Name
			}
		}

		if err := inst.Validate(); err != nil {
			return err
		}

		if err := validate(inst.Chunked, true, localLengths); err != nil {
			return err
		}

		if len(inst.Cases) > 0 {
			for _, cs := range inst.Cases {
				if err := validate(cs.Instructions, isChunked, localLengths); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// sequenceState tracks the ordering rules for a sequence of instructions. A <break> starts a new sequence.
type sequenceState struct {
	reachedOptional     bool
	reachedDummy        bool
	reachedUnsizedArray bool
	fieldNames          map[string]bool
	// switchField is the field of the switch in the current scope. It is kept across a break, since chunks share a scope.
	switchField string
}

func (st sequenceState) merge(other sequenceState) sequenceState {
	st.reachedOptional = st.reachedOptional || other.reachedOptional
	st.reachedDummy = st.reachedDummy || other.reachedDummy
	st.reachedUnsizedArray = st.reachedUnsizedArray || other.reachedUnsizedArray
	return st
}

func validateSequence(instructions []ProtocolInstruction, isChunked bool, state *sequenceState) error {
	for _, inst := range instructions {
		kind := inst.XMLName.Local

		if state.reachedDummy {
			return fmt.Errorf("validation error: <dummy> elements must not be followed by any other elements")
		}

		if state.reachedUnsizedArray && kind != "break" {
			return fmt.Errorf("validation error: non-delimited arrays without a length must be the final element (or the final element in the chunk, if chunked reading is enabled)")
		}

		switch kind {
		case "field", "array", "length":
			optional := inst.Optional != nil && *inst.Optional
			if state.reachedOptional && !optional {
				return fmt.Errorf("validation error: optional fields may not be followed by non-optional fields (%s)", instructionDescription(inst))
			}
			state.reachedOptional = state.reachedOptional || optional

			if inst.Name != nil {
				state.fieldNames[*inst.Name] = true
			}

			if kind == "array" {
				delimited := inst.Delimited != nil && *inst.Delimited
				if delimited && !isChunked {
					return fmt.Errorf("validation error: delimited arrays are only allowed in chunked sections (%s)", instructionDescription(inst))
				}
				if !delimited && inst.Length == nil {
					state.reachedUnsizedArray = true
				}
			}
		case "dummy":
			state.reachedDummy = true
		case "break":
			*state = sequenceState{fieldNames: state.fieldNames, switchField: state.switchField}
		case "chunked":
			if err := validateSequence(inst.Chunked, true, state); err != nil {
				return err
			}
		case "switch":
			if err := validateSwitch(inst, state); err != nil {
				return err
			}

			merged := *state
			for _, cs := range inst.Cases {
				caseState := *state
				caseState.switchField = ""
				if err := validateSequence(cs.Instructions, isChunked, &caseState); err != nil {
					return err
				}
				merged = merged.merge(caseState)
			}
			*state = merged
		}
	}

	return nil
}

func validateSwitch(inst ProtocolInstruction, state *sequenceState) error {
	if inst.Field == nil || !state.fieldNames[*inst.Field] {
		return fmt.Errorf("validation error: switch must reference a preceding field (%s)", instructionDescription(inst))
	}

	// switch factories set one switch per scope
	if len(state.switchField) > 0 {
		return fmt.Errorf("validation error: switch factories don't support multiple switches in one scope (switches on %s and %s)", state.switchField, *inst.Field)
	}
	state.switchField = *inst.Field

	values := map[string]bool{}
	for i, cs := range inst.Cases {
		if cs.Default {
			if i != len(inst.Cases)-1 {
				return fmt.Errorf("validation error: only the last case in a switch on %s can be the default case", *inst.Field)
			}
			continue
		}

		if values[cs.Value] {
			return fmt.Errorf("validation error: duplicate case value %s in switch on %s", cs.Value, *inst.Field)
		}
		values[cs.Value] = true
	}

	return nil
}

func instructionDescription(inst ProtocolInstruction) string {
	if inst.Name != nil {
		return fmt.Sprintf("<%s> %s", inst.XMLName.Local, *inst.Name)
	}
	if inst.Field != nil {
		return fmt.Sprintf("<%s> %s", inst.XMLName.Local, *inst.Field)
	}
	return fmt.Sprintf("<%s>", inst.XMLName.Local)
}

func validateInstructions(typeName string, instructions []ProtocolInstruction) error {
	if err := validate(instructions, false, nil); err != nil {
		return fmt.Errorf("%s: %w", typeName, err)
	}

	if err := validateSequence(instructions, false, &sequenceState{fieldNames: map[string]bool{}}); err != nil {
		return fmt.Errorf("%s: %w", typeName, err)
	}

	return nil
}

func (p Protocol) Validate() error {
	for _, st := range p.Structs {
		if err := validateInstructions(st.Name, st.Instructions); err != nil {
			return err
		}
	}

	for _, pkt := range p.Packets {
		if err := validateInstructions(pkt.GetTypeName(), pkt.Instructions); err != nil {
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

		if fieldName.Name == "XMLName" || fieldName.Name == "IsChunked" || fieldName.Name == "ReferencedBy" {
			continue
		}

		if utils.FindIndex(all, fieldName.Name, strEq) == -1 {
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

		if utils.FindIndex(required, fieldName.Name, strEq) >= 0 && fieldValue.IsNil() {
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
		all = []string{"Field", "Cases", "Comment"}
		required = []string{"Field", "Cases"}
	case "chunked":
		all = []string{"Chunked", "Comment"}
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
