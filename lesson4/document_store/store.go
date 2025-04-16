package document_store

type Store struct {
	collection map[string]*Collection //мапа по зберіганні колекцій по імені
}

// створення нового Store
func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
// Якщо ж колекція вже створеня то повертаємо `false` та nil
func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	if _, exists := s.Collection[name]; exists {
		return false, nil
	} // перевірка чи колекція існує, і якщо існує - вивиодимо false, nil

	newCollection := &Collection{
		Config:   *cfg,
		Document: make(map[string]Document),
	}
	s.Collection[name] = newCollection //додовання колекції у мапу
	return true, newCollection         //створено - виводимо True і нейм кол
}

// метод отримання колекції по назві/імені
func (s *Store) GetCollection(name string) (*Collection, bool) {
	col, found := s.Collection[name]
	if !found {
		return nil, false
	}
	return col, true
}

// видалення колекції
func (s *Store) DeleteCollection(name string) bool {
	_, found := s.Collection[name]
	if found {
		delete(s.Collection, name)
		return true
	}
	return false
}
