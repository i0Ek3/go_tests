package maps

import (
//"errors"
)

// ErrXxxx defines the error message
const (
	ErrNotFound = MapsErr("cannot find the value")
	ErrExist    = MapsErr("value already exists")
	ErrNotExist = MapsErr("value not exists")
)

// MapsErr defines type alias
type MapsErr string

// Error() defines the map error masssges
func (e MapsErr) Error() string {
	return string(e)
}

// Maps defines the map type
type Maps map[int]string

// Mapper as Maps's interface
type Mapper interface {
	Search(id int) string
	Add(id int, value string) error
	Update(id int, value string) error
	Delete(id int) error
}

// Search searchs value from map
func (m Maps) Search(id int) (string, error) {
	result, ok := m[id]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

// Add adds new item into map
func (m Maps) Add(id int, value string) error {
	//_, ok := m[id]
	//if !ok {
	//	m[id] = value
	//}
	//return ErrExist

	_, err := m.Search(id)
	switch err {
	case ErrNotFound:
		m[id] = value
	case nil:
		return ErrExist
	default:
		return err
	}

	return nil
}

// Update updates the value of map
func (m Maps) Update(id int, value string) error {
	_, err := m.Search(id)
	switch err {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		m[id] = value
	default:
		return err
	}

	return nil
}

// Delete deletes the value of map
// delete the not existing value we'll show you error
func (m Maps) Delete(id int) error {
	_, err := m.Search(id)
	switch err {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		delete(m, id)
	default:
		return err
	}

	return nil
}
