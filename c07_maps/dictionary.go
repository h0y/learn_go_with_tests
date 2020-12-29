package maps

import "errors"

type Dictionary map[string]string

var NotFoundErr = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", NotFoundErr
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}