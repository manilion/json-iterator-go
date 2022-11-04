package extra

import (
	"strings"
	"unicode"

	"github.com/manilion/json-iterator-go"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&privateFieldsExtension{})
}

type privateFieldsExtension struct {
	jsoniter.DummyExtension
}

func (extension *privateFieldsExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, binding := range structDescriptor.Fields {
		isPrivate := unicode.IsLower(rune(binding.Field.Name()[0]))
		if isPrivate {
			tag, hastag := binding.Field.Tag().Lookup("json")
			if !hastag {
				continue
			}
			tagParts := strings.Split(tag, ",")
			names := calcFieldNames(binding.Field.Name(), tagParts[0], tag)
			binding.FromNames = names
			binding.ToNames = names
		}
	}
}

func calcFieldNames(originalFieldName string, tagProvidedFieldName string, wholeTag string) []string {
	// ignore?
	if wholeTag == "-" {
		return []string{}
	}
	// rename?
	var fieldNames []string
	if tagProvidedFieldName == "" {
		fieldNames = []string{originalFieldName}
	} else {
		fieldNames = []string{tagProvidedFieldName}
	}
	return fieldNames
}
