package document_store

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string" // для полів типу строка.
	DocumentFieldTypeNumber DocumentFieldType = "number" // для полів типу число
	DocumentFieldTypeBool   DocumentFieldType = "bool"   // для полів true/false
	DocumentFieldTypeArray  DocumentFieldType = "array"  // для полів типу масив
	DocumentFieldTypeObject DocumentFieldType = "object" // для полів об'єкт
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}
