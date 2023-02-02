package internal

type DictEntry struct {
	Key  *Object
	Val  *Object
	Next *DictEntry
}
