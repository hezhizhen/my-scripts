package util

import (
	"fmt"
	"reflect"
)

// OrderedMap is like a simple map but stores keys' in order as well.
type OrderedMap struct {
	hmap  map[interface{}]interface{}
	slice []interface{}
}

// New creates a pointer pointing to an ordered map.
func New() *OrderedMap {
	return &OrderedMap{
		hmap: make(map[interface{}]interface{}),
	}
}

// Set adds or updates the record of given key.
func (m *OrderedMap) Set(key, value interface{}) error {
	if err := checkType(m.slice, key); err != nil {
		return err
	}
	if _, exist := m.hmap[key]; !exist {
		m.slice = append(m.slice, key)
	}
	// add or update
	m.hmap[key] = value
	return nil
}

// Delete deletes the record of given key.
func (m *OrderedMap) Delete(key interface{}) error {
	if err := checkType(m.slice, key); err != nil {
		return err
	}
	if _, exist := m.hmap[key]; !exist {
		return nil
	}
	delete(m.hmap, key)
	var newSlice []interface{}
	for _, item := range m.slice {
		if item == key {
			continue
		}
		newSlice = append(newSlice, item)
	}
	m.slice = newSlice
	return nil
}

// Pretty prints each key-value pair in one line.
func (m *OrderedMap) Pretty() {
	for _, key := range m.slice {
		fmt.Printf("%v:%v\n", key, m.hmap[key])
	}
}

func checkType(slice []interface{}, item interface{}) error {
	// TODO: can't check element's real type of slice
	elemT := reflect.TypeOf(slice).Elem().Kind()
	itemT := reflect.ValueOf(item).Kind()
	if elemT != itemT {
		return fmt.Errorf("slice element (%v) and given item (%v) have different types", elemT, itemT)
	}
	return nil
}
