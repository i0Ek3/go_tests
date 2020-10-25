package maps

import (
	"errors"
)

// ErrNotFound defines the error message
var ErrNotFound = errors.New("cannot find the value")

// Maps defines the map type
type Maps map[int]string

// Mapper as Maps's interface
type Mapper interface {
	Search(id int) string
}

// Search searchs value from map
func (m Maps) Search(id int) (string, error) {
	result, ok := m[id]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}
