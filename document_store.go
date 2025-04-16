package document_store

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string" // для полів типу строка.
	DocumentFieldTypeNumber DocumentFieldType = "number" // для полів типу число
	DocumentFieldTypeBool   DocumentFieldType = "bool"   // для полів true/false
	DocumentFieldTypeArray  DocumentFieldType = "array"  // для полів типу масив
	DocumentFieldTypeObject DocumentFieldType = "object" // для полів об'єкт
)

var ErrDocumentNotFound = errors.New("document not found")
var ErrInvalidDocumentInput = errors.New("invalid document input or output is nil")

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

func MarshalDocument(input any) (*Document, error) {
	if input == nil {
		return nil, ErrDocumentNotFound
	}
	jsonDoc, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var fieldsMap map[string]interface{}
	err = json.Unmarshal(jsonDoc, &fieldsMap)
	if err != nil {
		return nil, err
	} // перетворення JSON назад у map

	doc := &Document{
		Fields: make(map[string]DocumentField),
	}

	for key, value := range fieldsMap { //перевіряємо тип кожного поля за допомогою switch
		var fieldType DocumentFieldType

		switch value.(type) {
		case string:
			fieldType = DocumentFieldTypeString
		case float64:
			fieldType = DocumentFieldTypeNumber
		case bool:
			fieldType = DocumentFieldTypeBool
		case []interface{}:
			fieldType = DocumentFieldTypeArray
		case map[string]interface{}:
			fieldType = DocumentFieldTypeObject
		default:
			return nil, fmt.Errorf("%w: %s", key)
		}

		doc.Fields[key] = DocumentField{
			Type:  fieldType,
			Value: value,
		}
	}
	return doc, nil
}

func UnmarshalDocument(doc *Document, output any) error {
	if doc == nil || output == nil {
		return ErrInvalidDocumentInput //перевіряємо чи не порожній документ
	}

	tempMap := make(map[string]interface{}) //  тимчасова мапа для зберігання значень

	for key, field := range doc.Fields {
		tempMap[key] = field.Value //перенесення значень з док у мапу
	}
	jsonDoc, err := json.Marshal(tempMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonDoc, output)
	if err != nil {
		return err
	}
	return nil
}

type MyStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Block bool   `json:"block"`
}

func marshalExample() {
	s := &MyStruct{
		Name:  "Dania",
		Age:   20,
		Block: true,
	}

	doc, err := MarshalDocument(s)
	if err != nil {
		fmt.Printf("failed to marshal document: %v\n", err)
		return
	}

	fmt.Printf("marshaled document: %v\n", doc)
}

func unmarshalExample() {
	doc := &Document{
		Fields: map[string]DocumentField{
			"Name": {
				Type:  DocumentFieldTypeString,
				Value: "Dania",
			},
			"Age": {
				Type:  DocumentFieldTypeNumber,
				Value: 20,
			},
			"Block": {
				Type:  DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	var s MyStruct

	err := UnmarshalDocument(doc, &s)
	if err != nil {
		fmt.Printf("failed to unmarshal document: %v\n", err)
		return
	}

	fmt.Printf("unmarshaled document: %v\n", s)
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Block: %t\n", s.Block)
}

//func main() {
//	marshalExample()
//	fmt.Println("")
//	unmarshalExample()
//}
