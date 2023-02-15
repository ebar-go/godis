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

func (table *HashTable) Len() int {
	return len(table.items)
}

func (table *HashTable) Del(field string) {
	delete(table.items, field)
}

func (table *HashTable) Fields() []string {
	fields := make([]string, 0, len(table.items))
	for field, _ := range table.items {
		fields = append(fields, field)
	}
	return fields
}
