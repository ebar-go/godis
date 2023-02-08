package store

type Store struct {
	dict *Dict
}

func NewStore() *Store {
	return &Store{dict: NewDict()}
}
