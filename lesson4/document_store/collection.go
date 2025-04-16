package document_store

import "fmt"

// Структура Collection для зберігання док
type Collection struct {
	Config   CollectionConfig
	Document map[string]Document
} // мапа для зберігання документів по ключу

// структура конфігурації
type CollectionConfig struct {
	PrimaryKey string
}

// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`
func (s *Collection) Put(doc Document) {
	keyField, exists := doc.Fields[s.Config.PrimaryKey]
	if !exists {
		fmt.Println("keyField does not exist") // виводиться, якщо полене знайдено
		return
	}
	if keyField.Type != DocumentFieldTypeString {
		fmt.Println("keyField is not string")
	} // перевірка на тип STRING

	s.Document[keyField.Value.(string)] = doc // якщо все підходить - додається в колекцію
}

// отримання док з колекції за допомогою ключа
func (s *Collection) Get(key string) (*Document, bool) {
	doc, found := s.Document[key]
	return &doc, found
	if !found {
		return nil, false
	}
	return &doc, true
}

// видалення документа з колекції по ключу
func (s *Collection) Delete(key string) bool {
	_, found := s.Document[key]
	if found {
		delete(s.Document, key)
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	result := []Document{}
	for _, doc := range s.Document {
		result = append(result, doc)
	}
	return result
}
