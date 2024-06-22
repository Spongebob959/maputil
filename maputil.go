package maputil

import (
	"fmt"
)

type MapUtility[K comparable, V any] interface {
	CreateMap(keys []K, values []V) (map[K]V, error)
	MapToSlice(m map[K]V) ([]K, []V)
	GetKeyByIndex(index int, m map[K]V) K
	GetValByIndex(index int, m map[K]V) V
}

type MapUtilityImpl[K comparable, V any] struct{}

func (mu MapUtilityImpl[K, V]) CreateMap(keys []K, values []V) (map[K]V, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("amount of keys and values must be a straight number")
	}
	m := make(map[K]V)
	for i := 0; i < len(keys); i++ {
		m[keys[i]] = values[i]
	}
	return m, nil
}

func (mu MapUtilityImpl[K, V]) MapToSlices(m map[K]V) ([]K, []V) {
	keys := []K{}
	values := []V{}
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func (mu MapUtilityImpl[K, V]) GetKeyByIndex(index int, m map[K]V) (K, error) {
	s := []K{}
	for k := range m {
		s = append(s, k)
	}
	if index < 0 || index >= len(s) {
		var zero K
		return zero, fmt.Errorf("index out of range")
	}
	return s[index], nil
}

func (mu MapUtilityImpl[K, V]) GetValByIndex(index int, m map[K]V) (V, error) {
	key, err := mu.GetKeyByIndex(index, m)
	if err != nil {
		var zero V
		return zero, err
	}
	value := m[key]
	return value, nil
}

func (mu MapUtilityImpl[K, V]) PopByIndex(index int, m map[K]V) (error) {
	key, err := mu.GetKeyByIndex(index, m)
	if err != nil {
		return err
	}
	delete(m, key)
	return nil
}
