package types

type HashTable struct {
	items map[string]any
}

func NewHashTable() *HashTable {
	return &HashTable{items: map[string]any{}}
}

func (table *HashTable) Set(field string, value any) {
	table.items[field] = value
}

func (table *HashTable) Get(field string) any {
	return table.items[field]
}

func (table *HashTable) Has(field string) bool {
	_, ok := table.items[field]
	return ok
}
